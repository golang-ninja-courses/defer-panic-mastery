package main

var r = func() { recover() }

func main() {
	defer r()
	panic("sky is falling")
}
