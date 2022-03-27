package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Мы запускаем n горутин и через разное время убиваем
// каждую из них с помощью runtime.Goexit.

func main() {
	var wg sync.WaitGroup

	for i, sleep := range []time.Duration{
		150 * time.Millisecond,
		100 * time.Millisecond,
		50 * time.Millisecond,
	} {
		wg.Add(1)
		go func(i int, sleep time.Duration) {
			defer func() {
				wg.Done()
				fmt.Printf("recovered #%d: %v\n", i, recover()) // <- Всегда будет <nil>.
			}()

			worker(i, sleep)
		}(i, sleep)
	}

	wg.Wait()
	fmt.Println("DONE")
}

func worker(i int, sleep time.Duration) {
	defer fmt.Printf("worker #%d was killed\n", i)

	time.Sleep(sleep)
	runtime.Goexit() // <- Убиваем воркер, без влияния на остальные воркеры.

	fmt.Println("after work") // <- Этого мы не увидим, unreachable code.
}
