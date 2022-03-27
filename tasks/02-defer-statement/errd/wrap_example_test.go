package errd

import (
	"fmt"
	"io"
	"os"
)

func ExampleWrap() {
	err := Copy("unknown.txt", "known.txt")
	fmt.Println(err)

	// Output:
	// copy "unknown.txt" to "known.txt": open src file: open unknown.txt: no such file or directory
}

func Copy(src, dst string) error {
	return copyFile(src, dst)
}

func copyFile(src, dst string) (err error) {
	defer Wrap(&err, "copy %q to %q", src, dst)

	r, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open src file: %w", err)
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create dst file: %w", err)
	}
	defer w.Close()

	if _, err := io.Copy(w, r); err != nil {
		return err
	}
	return w.Sync()
}
