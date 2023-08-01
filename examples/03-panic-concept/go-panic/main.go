package main

func main() {
	defer func() { recover() }()
	go panic("crash me")
	select {}
}
