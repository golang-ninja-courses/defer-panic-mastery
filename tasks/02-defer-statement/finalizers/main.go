package main

import (
	"fmt"
	"runtime"
	"time"
)

type Hero struct {
	Name string
}

func NewHero(name string) *Hero {
	h := &Hero{Name: name}
	runtime.SetFinalizer(h, fin)
	return h
}

func fin(h *Hero) {
	fmt.Printf("%s was dead\n", h.Name)
}

func Battle() {
	n := NewHero("Naruto")
	_ = NewHero("Aoi Rokusho")

	go func() {
		for {
			runtime.KeepAlive(n)
		}
	}()
}

func main() {
	Battle()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}
