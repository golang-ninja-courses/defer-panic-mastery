package main

import "fmt"

func main() {
	var deferBits uint8
	printBits(deferBits)

	deferBits |= 1 << 0
	printBits(deferBits)

	deferBits |= 1 << 1
	printBits(deferBits)

	fmt.Println("Interrupt.")
	printBits(deferBits)

	deferBits &^= 1 << 1
	printBits(deferBits)

	deferBits &^= 1 << 0
	printBits(deferBits)
}

func printBits(v interface{}) {
	fmt.Printf("%08b\n", v)
}
