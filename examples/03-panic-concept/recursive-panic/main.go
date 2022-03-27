package main

func main() {
	alarm()
}

func alarm() {
	defer func() {
		recover()
		alarm()
	}()
	panic("alarm!")
}
