package filecache

import "fmt"

type Cache struct{}

func New() (*Cache, error) {
	return new(Cache), nil
}

func (c *Cache) CleanUp() error {
	fmt.Println("filecache cleaned out")
	return nil
}
