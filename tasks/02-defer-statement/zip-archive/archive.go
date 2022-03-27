package ziparchive

import (
	"errors"
	"io"
	"time"
)

var ErrNothingToArchive = errors.New("nothing to archive")

type File interface {
	io.ReadWriteCloser
	io.ReaderAt
	Stat() (FileInfo, error)
	Sync() error
}

type FS interface {
	Create(name string) (File, error)
	Open(name string) (File, error)
	Stat(name string) (FileInfo, error)
}

type FileInfo interface {
	Name() string
	Size() int64
	ModTime() time.Time
}

// Archive объединяет файлы по путям inPaths в один ZIP-архив по пути outPath.
// Если список входящих путей пуст, то возвращает ошибку ErrNothingToArchive.
func Archive(fSys FS, outPath string, inPaths ...string) error {
	// Реализуй меня.
	return nil
}
