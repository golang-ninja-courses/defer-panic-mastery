package main

import "testing"

func Benchmark_with9Defers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		with9Defers()
	}
}

func Benchmark_with9DefersHack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		with9DefersHack()
	}
}
