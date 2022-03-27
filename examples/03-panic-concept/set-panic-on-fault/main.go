package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"unsafe"
)

const kernelSpaceAddr = uintptr(0xffffffffffffffff)

func main() {
	defer handleFault()

	debug.SetPanicOnFault(true)
	// debug.SetPanicOnFault(false)

	v := *(*byte)(unsafe.Pointer(kernelSpaceAddr)) // Пытаемся залезть туда, куда нам нельзя.
	fmt.Println(v)
}

func handleFault() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)

		var addressable interface {
			Addr() uintptr
		}
		if err, ok := r.(error); ok && errors.As(err, &addressable) {
			fmt.Println("panic from addr:", addressable.Addr())
		}
	}
}
