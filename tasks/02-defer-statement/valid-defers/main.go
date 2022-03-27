package main

import (
	"sync"
)

type Service struct{}            // ...
func (s *Service) StaticMethod() {}

func main() {
	var (
		mu   sync.Mutex
		nums []int
		svc  Service
	)

	/*
		defer float64(100)
		defer (mu.Unlock())
		defer { result = float64(100) }
		defer func() { fmt.Println(`¯\_(ツ)_/¯`) }
		defer (go worker())
		defer func() { go worker() }()
	*/

	defer (*Service).StaticMethod(&svc)
	defer func() { nums = append(nums, 10) }()
	defer func() { go worker() }()
	defer worker()
	defer mu.Unlock()
}

func worker() {}
