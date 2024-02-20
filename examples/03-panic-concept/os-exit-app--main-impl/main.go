package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/golang-ninja-courses/defer-panic-mastery/examples/03-panic-concept/os-exit-app--resource-leak/app"
	"github.com/golang-ninja-courses/defer-panic-mastery/examples/03-panic-concept/os-exit-app--resource-leak/filecache"
	"github.com/golang-ninja-courses/defer-panic-mastery/examples/03-panic-concept/os-exit-app--resource-leak/store"
	"github.com/golang-ninja-courses/defer-panic-mastery/examples/03-panic-concept/os-exit-app--resource-leak/worker"
)

func main() {
	if err := mainImpl(); err != nil {
		log.Fatal(err)
	}
}

//nolint:errcheck
func mainImpl() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Запускаем фоновые задачи от корневого контекста.
	wrk, err := worker.New()
	if err != nil {
		return fmt.Errorf("build worker: %v", err)
	}
	defer wrk.Wait()
	go wrk.Run(ctx)

	// Конструируем приложение.
	cache, err := filecache.New()
	if err != nil {
		return fmt.Errorf("build cache: %v", err)
	}
	defer cache.CleanUp()

	storage, err := store.New()
	if err != nil {
		return fmt.Errorf("build storage: %v", err)
	}
	defer storage.Close()

	service, err := app.New(storage, cache)
	if err != nil {
		return fmt.Errorf("build app: %v", err)
	}

	// Запускаем приложение.
	return service.Start(ctx)
}
