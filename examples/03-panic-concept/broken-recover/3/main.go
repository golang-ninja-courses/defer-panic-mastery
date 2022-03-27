package main

import "fmt"

func main() {
	say("hello!", 3)
}

func say(s string, n int) {
	defer func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("got panic:", err)
			}
		}()

		cleanUp()
	}()
	for i := 0; i < n; i++ {
		if i == 1 {
			panic("keep silence")
		}
		fmt.Println(s)
	}
}

func cleanUp() {
	panic("while cleanup")
}
