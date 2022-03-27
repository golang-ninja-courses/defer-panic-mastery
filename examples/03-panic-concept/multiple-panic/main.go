package main

import (
	"fmt"
)

func main() {
	defer func() {
		defer func() {
			fmt.Println("before third panic")
			panic("triple kill")
		}()

		fmt.Println("before second panic")
		panic("double punch")
	}()

	fmt.Println("before first panic")
	panic("sky is falling")
}
