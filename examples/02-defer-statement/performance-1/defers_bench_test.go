package main

import "testing"

// https://go.dev/doc/manage-install#installing-multiple
// go1.12 test -benchmem -bench .
//
// go1.12 test -cpuprofile defers.out -bench .
// go tool pprof defers.out

func Benchmark_withoutDefer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withoutDefer()
	}
}

func Benchmark_withDefer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withDefer()
	}
}

func Benchmark_withDefers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withDefers()
	}
}
