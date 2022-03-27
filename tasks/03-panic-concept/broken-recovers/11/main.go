package main

import "fmt"

func main() {
	a := []int{99, 1, 42}

	var total int
	for i := 0; i <= len(a); i++ {
		total += a[i]
	}
	fmt.Println(total)
}

/*
defer func() {
	// Global recovering.
	recover()
}()
*/
