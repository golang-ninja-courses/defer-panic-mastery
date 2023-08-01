package main

import "fmt"

func main() {
	defer fmt.Println("defer 1")
	defer handlePanic()
	defer fmt.Println("defer 2")

	m := map[string]int{}
	go func() {
		defer fmt.Println("defer 3")
		defer handlePanic()
		defer fmt.Println("defer 4")
		for {
			m["x"] = 42
		}
	}()

	for {
		_ = m["x"]
	}
}

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("recovered:", err)
	}
}
