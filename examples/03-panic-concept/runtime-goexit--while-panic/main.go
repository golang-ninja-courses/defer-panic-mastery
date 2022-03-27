package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan struct{})
	go func() {
		defer close(c)
		defer runtime.Goexit()
		panic("boom!")
	}()

	<-c
	fmt.Println("OK")
}
