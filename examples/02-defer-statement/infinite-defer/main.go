package main

func main() {
	for {
		defer foo() //nolint:staticcheck
	}
}

func foo() {}
