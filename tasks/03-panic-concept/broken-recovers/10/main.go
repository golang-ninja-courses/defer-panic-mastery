package main

func main() {
	panic("sky is falling")
	recover() //nolint:govet
}
