package main

import "fmt"

func main() {
	fmt.Println(*calculate())
}

func calculate() *int {
	i := 3

	defer func() {
		i = 42
	}()

	i = 100
	return &i
}
