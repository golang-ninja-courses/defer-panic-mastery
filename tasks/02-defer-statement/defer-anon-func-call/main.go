package main

import "fmt"

func main() {
	s := map[int]string{100: "foo"}

	defer func() {
		fmt.Println("v =", s[100]+"_deferred")
	}()
	s[100] = "bar"
}
