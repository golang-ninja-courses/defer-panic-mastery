package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("deferred operation") // <- Это мы увидим.

	// Все горутинки полностью и успешно завершат
	// свою работу несмотря на Goexit ниже.
	go worker()
	go worker()
	go worker()
	runtime.Goexit()

	fmt.Println("hello!") // <- Это мы не увидим, unreachable code.
}

func worker() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("work: ", i)
	}
}
