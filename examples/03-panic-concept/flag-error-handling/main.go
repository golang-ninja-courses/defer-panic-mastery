package main

import (
	"flag"
	"fmt"
	"os"
)

// go run main.go -v
// go run main.go -k

func main() {
	debug := flag.NewFlagSet("debug", flag.PanicOnError) // flag.ContinueOnError, flag.ExitOnError, flag.PanicOnError
	v := debug.Bool("v", false, `Print app version.`)

	if err := debug.Parse(os.Args[1:]); err != nil {
		fmt.Println("Parse error:", err)
		return
	}
	if *v {
		fmt.Println("1.0.0")
	}
}
