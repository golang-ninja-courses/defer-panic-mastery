package store

import "fmt"

type Store struct{}

func New() (*Store, error) {
	return new(Store), nil
}

func (s *Store) Close() error {
	fmt.Println("storage connection closed")
	return nil
}
