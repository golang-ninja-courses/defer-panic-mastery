package main

import "fmt"

func main() {
	defer func() {
		cleanUp()
	}()
	fmt.Println("I'm OK")
}

func cleanUp() {
	defer func() {
		recover()
	}()

	panic("sky is falling - 2")
}
