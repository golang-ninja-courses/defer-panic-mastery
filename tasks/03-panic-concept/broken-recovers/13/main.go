package main

func init() {
	panic("sky is falling")
}

func init() {
	defer func() {
		recover()
	}()
}

func main() {
	// Go!
}
