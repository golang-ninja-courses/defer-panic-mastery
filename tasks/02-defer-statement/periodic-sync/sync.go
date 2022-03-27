package filehelpers

import (
	"context"
	"time"
)

type Syncer interface {
	Sync() error
}

// Sync синхронизирует входной s через интервалы времени, равные period.
// Является блокирующей функцией.
func Sync(ctx context.Context, s Syncer, period time.Duration) error {
	// Реализуй меня.
	return nil
}
