package main

import "fmt"

func main() {
	defer fmt.Println("1")
	defer func() {
		fmt.Println("2")
		fmt.Println("3")
	}()

	{
		defer fmt.Println("SCOPE")
	}

	for i := 11; i <= 13; i++ {
		defer fmt.Println(i)
		fmt.Println(i)
	}

	fmt.Println("RETURN")
}
