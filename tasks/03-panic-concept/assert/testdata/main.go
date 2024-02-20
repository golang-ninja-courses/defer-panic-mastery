package main

import (
	"fmt"

	"github.com/golang-ninja-courses/defer-panic-mastery/tasks/03-panic-concept/assert"
)

func main() {
	var v *int
	assert.Assert(v != nil, "v must be initialized")

	fmt.Println("I'm OK")
}
