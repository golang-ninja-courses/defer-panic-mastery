package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main")

	go func() {
		defer fmt.Println("go 1")
		time.Sleep(time.Second) // Задержка нужна, чтобы горутины были живы на момент выхода из main.
	}()

	go func() {
		defer fmt.Println("go 2")
		time.Sleep(time.Second)
	}()

	defer fmt.Println("before panic")
	{
		panic("sky is falling")
	}
	defer fmt.Println("after panic") //nolint:govet
}
