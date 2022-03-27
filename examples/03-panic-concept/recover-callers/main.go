package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		printFrames("func deferred in main", frames(2))
		if r := recover(); r != nil {
			fmt.Println("\nrecovered", r)
		}
	}()

	bar()
	fmt.Println("don't printed")
}

//go:noinline
func bar() {
	foo()
}

//go:noinline
func foo() {
	printFrames("foo", frames(2))
	panic("sky is falling")
}

func frames(skip int) *runtime.Frames {
	pcs := make([]uintptr, 8)
	return runtime.CallersFrames(pcs[:runtime.Callers(skip, pcs)])
}

func printFrames(from string, f *runtime.Frames) {
	fmt.Printf("\nFrom %q:\n", from)
	for {
		frame, more := f.Next()
		fmt.Println(frame.PC, frame.Function)
		if !more {
			break
		}
	}
}
