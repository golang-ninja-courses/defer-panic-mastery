package main

import "fmt"

func main() {
	perfEventSysfsInitV1()
	fmt.Println("goto - OK")

	perfEventSysfsInitV2()
	fmt.Println("inplace - OK")

	perfEventSysfsInitV3()
	fmt.Println("defer - OK")
}
