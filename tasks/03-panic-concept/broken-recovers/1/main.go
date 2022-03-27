package main

import "fmt"

func main() {
	defer recover()
	foo()
	fmt.Println("I'm OK")
}

func foo() {
	panic("sky is falling")
}
