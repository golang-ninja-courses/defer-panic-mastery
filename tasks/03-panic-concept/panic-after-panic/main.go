package main

import "fmt"

func main() {
	process()
	fmt.Println("OK")
}

func process() {
	defer func() {
		recover()
	}()

	panic(1)
	panic(2) //nolint:govet
}
