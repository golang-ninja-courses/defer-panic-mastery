package main

func init() {
	defer func() {
		recover()
	}()
}

func init() {
	panic("sky is falling")
}

func main() {
	// Go!
}
