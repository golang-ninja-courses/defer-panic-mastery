package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("HELLO!") // Отработает последней.
	}()
	fmt.Println("WORLD") // Отработает первой.
}

/*
WORLD
HELLO!
*/
