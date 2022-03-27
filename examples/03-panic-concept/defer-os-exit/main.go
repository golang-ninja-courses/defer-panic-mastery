package main

import (
	"log"
	"os"
)

func main() {
	var exitCode int
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
		os.Exit(exitCode)
	}()

	if err = parseConfig(); err != nil {
		exitCode = 2
		return
	}

	if err = run(); err != nil {
		exitCode = 1
		return
	}
}

func parseConfig() error { return nil }
func run() error         { return nil }
