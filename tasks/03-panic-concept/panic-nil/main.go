package main

import "fmt"

func main() {
	doWork()
	fmt.Println("OK")
}

func doWork() {
	panic(nil)
}
