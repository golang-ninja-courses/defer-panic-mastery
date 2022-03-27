package main

import (
	"fmt"

	"github.com/www-golang-courses-ru/advanced-dealing-with-panic-in-go/tasks/03-panic-concept/assert"
)

func main() {
	var v *int
	assert.Assert(v != nil, "v must be initialized")

	fmt.Println("I'm OK")
}
