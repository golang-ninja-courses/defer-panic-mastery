package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	files := make([]*os.File, 3)
	_ = Upload(files)
}

func Upload(files []*os.File) error {
	var wg sync.WaitGroup

	for i, f := range files {
		i, f := i, f

		wg.Add(1)
		safeGo(func() {
			// defer wg.Done()
			upload(i, f)
			wg.Done()
		})
	}

	wg.Wait()
	return nil
}

func upload(i int, _ *os.File) {
	if i == 1 {
		_ = ([]int{})[0]
	}
	fmt.Printf("file #%d uploaded\n", i)
}

func safeGo(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recovered:", err)
		}
	}()
	f()
}
