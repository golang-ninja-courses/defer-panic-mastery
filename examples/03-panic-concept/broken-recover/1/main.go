package main

import "fmt"

func main() {
	defer recover() // <- Не сработает, recover() слишком "близко".

	v := calculate(1, 0)
	fmt.Println(v)
}

func calculate(alpha, n int) int {
	if n == 0 {
		panic(`"n" cannot be 0 due to algo`)
	}
	return alpha / n
}
