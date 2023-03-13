package ziparchive

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
	"testing/fstest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArchive_SmokeWithRealFS(t *testing.T) {
	// NOTE: Don't use filepath.Join
	// https://stepik.org/lesson/590501/step/4?discussion=7302903&unit=585453
	inPaths := []string{
		"testdata/1.txt",
		"testdata/2.txt",
		"testdata/3.txt",
	}
	outPath := "proverbs.zip"

	fSys := newFSMock()
	for _, p := range inPaths {
		require.NoError(t, fSys.Load(p))
	}

	err := Archive(fSys, outPath, inPaths...)
	require.NoError(t, err)

	err = fSys.Dump(outPath)
	require.NoError(t, err)

	z, err := zip.OpenReader(outPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, z.Close()) }()

	err = fstest.TestFS(z, inPaths...)
	require.NoError(t, err)

	// $ unzip -l proverbs.zip
	// Archive:  proverbs.zip
	//  Length      Date    Time    Name
	// ---------  ---------- -----   ----
	//	 249  01-30-2022 15:36   testdata/1.txt
	//	 279  01-30-2022 15:37   testdata/2.txt
	//	 233  01-30-2022 15:38   testdata/3.txt
	// ---------                     -------
	//	 761                     3 files
}

func TestArchive(t *testing.T) {
	fSys := newFSMock()

	d1 := time.Date(2022, 1, 1, 1, 1, 1, 0, time.UTC)
	d2 := time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC)

	fSys.files["hello.txt"] = &fileMock{
		data: bytes.NewBuffer([]byte(`Hello`)),
		info: fileInfoMock{name: "hello.txt", size: 5, modTime: d1},
	}
	fSys.files["world.txt"] = &fileMock{
		data: bytes.NewBuffer([]byte(`World!`)),
		info: fileInfoMock{name: "world.txt", size: 6, modTime: d2},
	}

	const outPath = "hello_world.zip"
	err := Archive(fSys, outPath, "hello.txt", "world.txt")
	require.NoError(t, err)

	{
		for _, f := range fSys.files {
			assert.True(t, f.closed, "file %q is not closed", f.info.Name())
		}

		f, ok := fSys.files[outPath]
		require.True(t, ok, "result archive %q is not created", outPath)
		assert.True(t, f.synced, "result archive %q is not synced", outPath)
	}

	f, err := fSys.Open(outPath)
	require.NoError(t, err)
	defer f.Close()

	s, err := f.Stat()
	require.NoError(t, err)

	zr, err := zip.NewReader(f, s.Size())
	require.NoError(t, err)
	require.Len(t, zr.File, 2)

	{
		zf1 := zr.File[0]
		assert.Equal(t, "hello.txt", zf1.Name)
		assert.Equal(t, uint64(5), zf1.UncompressedSize64)
		assert.Equal(t, d1, zf1.Modified.UTC())

		f1, err := zf1.Open()
		require.NoError(t, err)
		defer f1.Close()
		assertFileContent(t, f1, "Hello")
	}

	{
		zf2 := zr.File[1]
		assert.Equal(t, "world.txt", zf2.Name)
		assert.Equal(t, uint64(6), zf2.UncompressedSize64)
		assert.Equal(t, d2, zf2.Modified.UTC())

		f2, err := zf2.Open()
		require.NoError(t, err)
		defer f2.Close()
		assertFileContent(t, f2, "World!")
	}
}

func assertFileContent(t *testing.T, f io.Reader, expected string) {
	t.Helper()

	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)
	assert.Equal(t, expected, string(data))
}

func TestArchive_NoInput(t *testing.T) {
	for i, inputs := range [][]string{nil, {}} {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			fSys := newFSMock()
			err := Archive(fSys, "out.txt", inputs...)
			assert.ErrorIs(t, err, ErrNothingToArchive)
			assert.Empty(t, fSys.files)
		})
	}
}
