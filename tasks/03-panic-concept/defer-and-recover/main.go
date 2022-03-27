package main

import "fmt"

func main() {
	defer fmt.Println("before recover")
	{
		defer func() { recover() }()
	}
	defer fmt.Println("after recover")

	defer fmt.Println("before panic")
	{
		panic("sky is falling")
	}
	defer fmt.Println("after panic") //nolint:govet
}
