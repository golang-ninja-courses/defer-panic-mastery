package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/golang-ninja-courses/defer-panic-mastery/examples/03-panic-concept/os-exit-app--resource-leak/app"
	"github.com/golang-ninja-courses/defer-panic-mastery/examples/03-panic-concept/os-exit-app--resource-leak/filecache"
	"github.com/golang-ninja-courses/defer-panic-mastery/examples/03-panic-concept/os-exit-app--resource-leak/store"
	"github.com/golang-ninja-courses/defer-panic-mastery/examples/03-panic-concept/os-exit-app--resource-leak/worker"
)

//nolint:errcheck
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Запускаем фоновые задачи от корневого контекста.
	wrk, err := worker.New()
	mustNil(err, "build worker")
	defer wrk.Wait()
	go wrk.Run(ctx)

	// Конструируем приложение.
	cache, err := filecache.New()
	mustNil(err, "build cache")
	defer cache.CleanUp()

	storage, err := store.New()
	mustNil(err, "build storage")
	defer storage.Close()

	service, err := app.New(storage, cache)
	mustNil(err, "build app")

	// Запускаем приложение.
	err = service.Start(ctx)
	mustNil(err, "start service")
}

func mustNil(err error, msg string) {
	if err != nil {
		log.Panic(msg + ": " + err.Error())
	}
}
