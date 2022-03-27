package main

import "fmt"

func main() {
	fmt.Println(calculate())
}

func calculate() int {
	i := 3

	defer changeResult(&i)

	i = 100
	return i
}

func changeResult(i *int) {
	*i = 42
}
