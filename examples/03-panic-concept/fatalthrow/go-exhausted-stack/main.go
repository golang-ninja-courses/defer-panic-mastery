package main

import "fmt"

type fatArray [1 << 20]int64

func main() {
	defer handlePanic()

	var f func(a fatArray)
	f = func(a fatArray) {
		f(a)
	}
	f(fatArray{})
}

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("recovered:", err)
	}
}
