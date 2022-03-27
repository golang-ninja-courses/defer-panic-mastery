package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/www-golang-courses-ru/advanced-dealing-with-panic-in-go/examples/03-panic-concept/os-exit-app--resource-leak/app"
	"github.com/www-golang-courses-ru/advanced-dealing-with-panic-in-go/examples/03-panic-concept/os-exit-app--resource-leak/filecache"
	"github.com/www-golang-courses-ru/advanced-dealing-with-panic-in-go/examples/03-panic-concept/os-exit-app--resource-leak/store"
	"github.com/www-golang-courses-ru/advanced-dealing-with-panic-in-go/examples/03-panic-concept/os-exit-app--resource-leak/worker"
)

//nolint:errcheck
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Запускаем фоновые задачи от корневого контекста.
	wrk, err := worker.New()
	mustNil(err, "build worker") // <- Утечка defer'ов выше.
	defer wrk.Wait()
	go wrk.Run(ctx)

	// Конструируем приложение.
	cache, err := filecache.New()
	mustNil(err, "build cache") // <- Утечка defer'ов выше.
	defer cache.CleanUp()

	storage, err := store.New()
	mustNil(err, "build storage") // <- Утечка defer'ов выше.
	defer storage.Close()

	service, err := app.New(storage, cache)
	mustNil(err, "build app") // <- Утечка defer'ов выше.

	// Запускаем приложение.
	err = service.Start(ctx)
	mustNil(err, "start service")
}

func mustNil(err error, msg string) {
	if err != nil {
		log.Fatal(msg + ": " + err.Error())
	}
}
