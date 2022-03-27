package processor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ProcessFiles(paths []string) ([]string, error) {
	results := make([]string, 0, len(paths))

	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return nil, fmt.Errorf("open file %q: %v", p, err)
		}
		defer f.Close()

		if strings.HasSuffix(f.Name(), "_skip") {
			continue
		}

		r, err := processFile(f)
		if err != nil {
			return nil, fmt.Errorf("process file %q: %v", p, err)
		}

		results = append(results, r)
	}

	return results, nil
}

func processFile(f io.Reader) (string, error) {
	l, _, err := bufio.NewReader(f).ReadLine()
	if err != nil {
		return "", err
	}
	return string(l), nil
}
