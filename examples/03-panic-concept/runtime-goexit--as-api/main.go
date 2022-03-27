package main

import (
	"fmt"
	"runtime"
)

func main() {
	r := NewRunner()

	r.Run("first fn", func(r *R) {
		fmt.Println("first function is done")
	})

	r.Run("second fn", func(r *R) {
		defer fmt.Println("second function is done (defer)")

		r.Interrupt() // Внутри функции пользуемся API того, кто эту функцию будет запускать.
		fmt.Println("unreachable code")
	})

	fmt.Println()
	fmt.Println("first fn was interrupted?", r.WasInterrupted("first fn"))   // false
	fmt.Println("second fn was interrupted?", r.WasInterrupted("second fn")) // true
}

type fn = string

type Runner struct {
	interrupted map[fn]bool
}

func NewRunner() *Runner {
	return &Runner{
		interrupted: make(map[fn]bool),
	}
}

func (r *Runner) Run(name string, fn func(r *R)) {
	done := make(chan struct{})
	go func() {
		defer close(done)
		fn(&R{name: name, Runner: r})
	}()
	<-done
}

func (r *Runner) WasInterrupted(fnName string) bool {
	return r.interrupted[fnName]
}

type R struct {
	name string
	*Runner
}

func (r *R) Interrupt() {
	r.interrupted[r.name] = true
	runtime.Goexit()
}
