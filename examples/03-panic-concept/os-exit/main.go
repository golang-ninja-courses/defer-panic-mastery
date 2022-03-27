package main

import (
	"fmt"
	"os"
	"time"
)

// Мы запускаем n спящих горутин и убиваем процесс во время их сна.

func main() {
	defer fmt.Println("EXIT FROM main") // <- Это мы не увидим.

	const n = 3
	for i := 0; i < n; i++ {
		go worker(i)
	}

	go killer(1)

	// Если убрать блокировку, то это будет эквивалентно
	// вызову os.Interrupt(0) в конце функции main и горутины
	// выше даже не успеют запуститься.
	select {}
}

func worker(i int) {
	fmt.Printf("worker %d: start\n", i)
	time.Sleep(10 * time.Second)
	fmt.Printf("worker %d: stop\n", i) // <- Это мы не увидим.
}

func killer(code int) {
	defer fmt.Printf("EXIT FROM killer") // <- Это мы не увидим.

	time.Sleep(time.Second)
	fmt.Println("DEAD")
	os.Exit(code) //nolint:gocritic
}
