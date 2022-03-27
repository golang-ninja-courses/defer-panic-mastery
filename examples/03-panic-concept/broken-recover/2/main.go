package main

import "fmt"

func main() {
	say("hello!", 3)
}

func say(s string, n int) {
	defer func() {
		logPanic() // <- Не сработает, recover() слишком "далеко".
		cleanUp()
	}()
	for i := 0; i < n; i++ {
		if i == 1 {
			panic("keep silence")
		}
		fmt.Println(s)
	}
}

func logPanic() {
	if err := recover(); err != nil {
		fmt.Println("got panic:", err)
	}
}

func cleanUp() {
	/* ... */
}
