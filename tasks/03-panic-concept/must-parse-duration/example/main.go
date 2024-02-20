package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	envconfig "github.com/golang-ninja-courses/defer-panic-mastery/tasks/03-panic-concept/must-parse-duration"
)

var syncPeriod = envconfig.MustParseDuration("SYNC_PERIOD")

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	t := time.NewTicker(syncPeriod)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-t.C:
			log.Println("sync...")
		}
	}
}
