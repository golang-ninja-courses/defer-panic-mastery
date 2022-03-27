package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()

	*(*int)(nil) = 0
	fmt.Println("I'm OK")
}
