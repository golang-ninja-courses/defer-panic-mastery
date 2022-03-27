package main

import "fmt"

type InvalidArgumentError struct {
	fun, arg, reason string
}

func NewInvalidArgumentError(fun string, arg string, reason string) *InvalidArgumentError {
	return &InvalidArgumentError{fun: fun, arg: arg, reason: reason}
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("invalid arg %q of func %q: %v", e.arg, e.fun, e.reason)
}

func div(a, b int) int {
	if b == 0 {
		panic(NewInvalidArgumentError("div", "b", "division by zero"))
	}
	return a / b
}

func main() {
	fmt.Printf("div(42, 0) = %v\n", div(42, 0))
	fmt.Println("OK") // Не напечатается!
}
