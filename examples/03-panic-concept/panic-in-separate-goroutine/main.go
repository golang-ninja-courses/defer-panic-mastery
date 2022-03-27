package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()

	recovered := make(chan struct{})
	go foo(recovered)

	<-recovered
	fmt.Println("OK")
}

func foo(recovered chan<- struct{}) {
	defer func() {
		defer close(recovered)
		recover()
	}()
	panic("sky is falling")
}
