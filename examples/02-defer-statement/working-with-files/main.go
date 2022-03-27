package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "02-defer-statement__working-with-files*.txt")
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	if err := WriteToFile(f.Name(), "Hello!"); err != nil {
		log.Fatal(err)
	}

	if err := ReadNFromFile(f.Name(), 100, os.Stdout); err != nil {
		log.Fatal(err)
	}

	// Hello!
}

func ReadNFromFile(fname string, nn int, to io.Writer) error {
	f, err := os.Open(fname)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	data := make([]byte, nn)
	n, err := f.Read(data)
	if err != nil {
		return fmt.Errorf("read from file: %w", err)
	}

	_, err = to.Write(data[:n])
	return err
}

func WriteToFile(fname, text string) error {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		return fmt.Errorf("write to file: %w", err)
	}

	return f.Sync()
}
