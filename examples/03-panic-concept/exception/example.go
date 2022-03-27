package main

import "fmt"

func div(a, b int) int {
	return a / b
}

func main() {
	fmt.Printf("div(42, 0) = %v\n", div(42, 0))
	fmt.Println("OK") // Не напечатается!
}
