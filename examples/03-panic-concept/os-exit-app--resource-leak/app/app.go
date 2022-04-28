package app

import (
	"context"
	"errors"
)

type (
	IStorage any
	ICache   any
)

type App struct {
	storage IStorage
	cache   ICache
}

func New(storage IStorage, cache ICache) (*App, error) {
	return &App{
		storage: storage,
		cache:   cache,
	}, nil
}

func (a *App) Start(ctx context.Context) error {
	return errors.New("unexpected error")
}
