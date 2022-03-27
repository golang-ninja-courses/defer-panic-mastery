package main

import "fmt"

func main() {
	fmt.Println(*calculate())
}

func calculate() *int {
	i := new(int)
	*i = 3

	defer func() {
		i = intPtr(42)
		*i = 50
	}()

	*i = 100
	return i
}

func intPtr(i int) *int {
	return &i
}
