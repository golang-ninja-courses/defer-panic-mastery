package main

import "fmt"

const (
	MaxVersion = 26
	GenVersion = 20 // Current version.
	MinVersion = 21
)

const (
	_ = uint(MaxVersion - GenVersion) // uint(26-20) = uint(6)
	_ = uint(GenVersion - MinVersion) // uint(20-21) = uint(-1) = error
)

func main() {
	fmt.Println("Compiled?")
}
