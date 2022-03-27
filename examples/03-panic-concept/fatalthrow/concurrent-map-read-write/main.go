package main

import "fmt"

func main() {
	defer handlePanic()

	m := map[string]int{}
	go func() {
		defer handlePanic()
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
