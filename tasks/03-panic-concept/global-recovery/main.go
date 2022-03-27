package main

import (
	"log"
	"os"
)

func main() {
	defer recoverPanic()
	runApp()
}

func runApp() {
	panic("some error in application logic")
}

func recoverPanic() {
	if err := recover(); err != nil {
		log.Println("unexpected error:", err)
		os.Exit(1)
	}
}
