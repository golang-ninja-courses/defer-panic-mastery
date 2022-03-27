package portscanner

import (
	"context"
	"io"
	"time"
)

const (
	maxPortNumber = 65535
	workers       = 10
	connTimeout   = 10 * time.Millisecond
)

type Conn interface {
	io.Closer
}

type Dialer interface {
	DialContext(ctx context.Context, network, address string) (Conn, error)
}

// Scan сканирует hostname на предмет открытых TCP-портов.
func Scan(d Dialer, hostname string) []int {
	// Реализуй меня.
	return nil
}
