package main

import "fmt"

func main() {
	fmt.Println("a:", a()) // 100
	fmt.Println()

	fmt.Println("b:", b()) // 100, не 42!
	fmt.Println()

	fmt.Println("c:", c()) // 99
	fmt.Println()

	fmt.Println("d:", *d()) // 42
}

func a() int {
	i := 3
	defer func() {
		fmt.Println("defer in a: i =", i)
	}()

	i = 100
	return i
}

func b() int {
	i := 3
	defer func() {
		i = 42
		fmt.Println("defer in b: i =", i)
	}()

	return 100
}

func c() (i int) {
	i = 3
	defer func() {
		i = 42
		fmt.Println("defer in b: i =", i)
	}()

	i = 100
	return i
}

func d() *int {
	i := new(int)

	*i = 3
	defer func() {
		*i = 42
		fmt.Println("defer in d: i =", *i)
	}()

	*i = 100
	return i
}
