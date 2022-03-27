package main

import "fmt"

func main() {
	defer handlePanic()
	select {}
}

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("recovered:", err)
	}
}
