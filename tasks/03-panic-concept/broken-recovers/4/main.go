package main

func main() {
	defer func() {
		func() {
			recover()
		}()
	}()
	panic("sky is falling")
}
