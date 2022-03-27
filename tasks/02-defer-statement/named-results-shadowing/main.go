package main

import (
	"fmt"
)

func main() {
	if n, err := calculate(); err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println("OK", n)
	}
}

func calculate() (n int, err error) {
	defer func() {
		// Если в результате (возможно) сложных и множественных вычислений ниже
		// мы получили отрицательное число, то значит в алгоритме ошибка.
		// Далее по курсу мы узнаем, что для таких случаев лучше подходит паника,
		// но цель этого примера состоит в другом.
		if n < 0 {
			n, err := 0, fmt.Errorf("invalid formula realization: got negative number %d", n)
			_, _ = n, err
		}
	}()

	// ...
	n = -1
	return
}
