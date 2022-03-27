package main

import "fmt"

func main() {
	defer handlePanic()

	var f func()
	go f()
}

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("recovered:", err)
	}
}
