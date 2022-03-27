package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("three!") // 2
		fmt.Println("two!")   // 3
		fmt.Println("one!")   // 4
	}()
	fmt.Println("return") // 1
}

/*
return
three!
two!
one!
*/
