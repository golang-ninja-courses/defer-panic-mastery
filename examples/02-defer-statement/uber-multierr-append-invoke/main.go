package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"go.uber.org/multierr"
)

func main() {
	if err := processFile("/etc/hosts"); err != nil { // false
		fmt.Printf("%v\n", err)
	}

	if err := processCloserMock(); err != nil {
		fmt.Printf("%v\n", err) // close error
	}

	if err := processCloserMockWithError(); err != nil {
		fmt.Printf("%v\n", err) // process error; close error
	}
}

// processFile – пример штатного использования multierr.AppendInvoke.
// Обратите внимание, что multierr.AppendInvoke используется вместе с
// указателем на именованный возвращаемый аргумент err.
func processFile(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open file: %v", err)
	}
	defer multierr.AppendInvoke(&err, multierr.Close(f))

	scanner := bufio.NewScanner(f)
	defer multierr.AppendInvoke(&err, multierr.Invoke(scanner.Err))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}

type closerMock struct{}

func (c closerMock) Close() error {
	return errors.New("close error")
}

func processCloserMock() (err error) {
	defer multierr.AppendInvoke(&err, multierr.Close(closerMock{}))
	return nil
}

func processCloserMockWithError() (err error) {
	defer multierr.AppendInvoke(&err, multierr.Close(closerMock{}))
	return errors.New("process error")
}
