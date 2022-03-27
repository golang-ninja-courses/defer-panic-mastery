package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(GenerateUniqueID())
	}
}

func GenerateUniqueID() string {
	id := strconv.Itoa(rand.Int())
	mustNewID(id)
	knownIDs[id] = struct{}{}
	return id
}

func mustNewID(id string) {
	if _, ok := knownIDs[id]; ok {
		panic("not unique ID: " + id)
	}
}

var knownIDs = map[string]struct{}{}
