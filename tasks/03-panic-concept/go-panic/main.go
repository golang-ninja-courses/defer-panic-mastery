package main

import "fmt"

func main() {
	go panic("1")
	fmt.Println("I'm OK")
}
