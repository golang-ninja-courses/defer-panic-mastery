package main

import "fmt"

func main() {
	defer func() { fmt.Println(recover()) }()
	defer func() { fmt.Println(recover()) }()
	defer func() { fmt.Println(recover()) }()
	work()
}

func work() {
	panic("holiday")
}
