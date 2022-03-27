package worker

import (
	"context"
	"fmt"
)

type Worker struct{}

func New() (*Worker, error) {
	return new(Worker), nil
}

func (w *Worker) Run(ctx context.Context) error {
	fmt.Println("worker run")
	<-ctx.Done()
	return nil
}

func (w *Worker) Wait() error {
	fmt.Println("worker done")
	return nil
}
