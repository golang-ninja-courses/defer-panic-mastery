package dbhelpers

import (
	"context"
	"errors"
)

var ErrConnectionLost = errors.New("connection lost")

type Pinger interface {
	Ping(ctx context.Context) error
}

// Ping возвращает результат вызова p.Ping(ctx),
// или ErrConnectionLost в случае его паникования.
func Ping(ctx context.Context, p Pinger) error {
	// Реализуй меня.
	return nil
}
