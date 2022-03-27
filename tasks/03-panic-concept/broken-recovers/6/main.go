package main

func main() {
	defer recoverable()
	panic("sky is falling")
}

func recoverable() {
	defer func() {
		recover()
	}()
}
