package main

func main() {
	defer func() {
		recover()
	}()

	*(*int)(nil) = 0
	panic("sky is falling")
}

func recover() { //nolint:predeclared
}
