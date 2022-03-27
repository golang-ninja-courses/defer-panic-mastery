package main

import "fmt"

func main() {
	foo()
}

func foo() {
	var f func()
	defer f() // Аналогично `defer (func())(nil)()`.

	fmt.Println("hello")
	fmt.Println("world")
} // <- Паника указывает сюда.
