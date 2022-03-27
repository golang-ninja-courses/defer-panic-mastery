package ziparchive

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)

var _ FS = (*fsMock)(nil)

type fsMock struct {
	files map[string]*fileMock
}

func newFSMock() *fsMock {
	return &fsMock{files: make(map[string]*fileMock)}
}

func (f *fsMock) Create(name string) (File, error) {
	fl := newFileMock(fileInfoMock{
		name:    filepath.Base(name),
		size:    0,
		modTime: time.Now(),
	})
	f.files[name] = fl
	return fl, nil
}

func (f *fsMock) Open(name string) (File, error) {
	fl, ok := f.files[name]
	if !ok {
		return nil, errors.New("file not found")
	}

	fl.closed = false
	fl.synced = false
	return fl, nil
}

func (f *fsMock) Stat(name string) (FileInfo, error) {
	fl, ok := f.files[name]
	if !ok {
		return nil, errors.New("file not found")
	}
	return fl.Stat()
}

// Load загружает файл из "реальной" файловой системы.
func (f *fsMock) Load(name string) error {
	src, err := os.Open(name)
	if err != nil {
		return err
	}
	defer src.Close()

	stat, err := os.Stat(name)
	if err != nil {
		return err
	}

	fl := newFileMock(fileInfoMock{
		name:    stat.Name(),
		size:    stat.Size(),
		modTime: stat.ModTime(),
	})
	f.files[name] = fl

	_, err = io.Copy(fl.data, src)
	return err
}

// Dump выгружает файл в "реальную" файловую систему.
func (f *fsMock) Dump(name string) error {
	fl, ok := f.files[name]
	if !ok {
		return errors.New("file not found")
	}

	dst, err := os.Create(name)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, fl.data); err != nil {
		return err
	}
	return fl.Sync()
}

var _ File = (*fileMock)(nil)

type fileMock struct {
	data   *bytes.Buffer
	info   fileInfoMock
	synced bool
	closed bool
}

func newFileMock(i fileInfoMock) *fileMock {
	return &fileMock{
		data: bytes.NewBuffer(nil),
		info: i,
	}
}

func (f *fileMock) Stat() (FileInfo, error) {
	return f.info, nil
}

func (f *fileMock) Read(p []byte) (int, error) {
	return f.data.Read(p)
}

func (f *fileMock) ReadAt(p []byte, off int64) (int, error) {
	if off >= int64(f.data.Len()) {
		return 0, io.EOF
	}
	return bytes.NewBuffer(f.data.Bytes()[off:]).Read(p)
}

func (f *fileMock) Write(p []byte) (int, error) {
	n, err := f.data.Write(p)
	f.info.size += int64(n)
	return n, err
}

func (f *fileMock) Close() error {
	f.closed = true
	return nil
}

func (f *fileMock) Sync() error {
	f.synced = true
	return nil
}

var _ FileInfo = fileInfoMock{}

type fileInfoMock struct {
	name    string
	size    int64
	modTime time.Time
}

func (fi fileInfoMock) Name() string       { return fi.name }
func (fi fileInfoMock) Size() int64        { return fi.size }
func (fi fileInfoMock) ModTime() time.Time { return fi.modTime }
