package main

import "fmt"

func main() {
	defer fmt.Println("one!")   // 4)
	defer fmt.Println("two!")   // 3)
	defer fmt.Println("three!") // 2)
	fmt.Println("return")       // 1)
}

/*
return
three!
two!
one!
*/
