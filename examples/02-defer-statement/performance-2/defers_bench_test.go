package main

import "testing"

// https://go.dev/doc/manage-install#installing-multiple
// go1.12 test -benchmem -bench . > 1.12.txt
// go1.14 test -benchmem -bench . > 1.14.txt
// benchstat -alpha 1.1 1.12.txt 1.14.txt

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

func Benchmark_withNotOpenDefer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withNotOpenDefer()
	}
}

func Benchmark_with8Defers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		with8Defers()
	}
}

func Benchmark_with9Defers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		with9Defers()
	}
}
