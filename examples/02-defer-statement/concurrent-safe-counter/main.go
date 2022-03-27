package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.RWMutex
	i  int
}

func (s *Counter) Inc() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.i++
}

func (s *Counter) Value() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.i
}

// go run -race main.go.
func main() {
	s := new(Counter)

	const n = 100

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			s.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(s.Value())
}
