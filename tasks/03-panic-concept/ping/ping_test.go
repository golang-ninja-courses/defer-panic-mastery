package dbhelpers

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	errBadConn := errors.New("driver: bad connection")

	cases := []struct {
		name      string
		ping      pingerMock
		expErr    error
		expErrMsg string
	}{
		{
			name: "no error",
			ping: func(_ context.Context) error {
				return nil
			},
			expErr: nil,
		},
		{
			name: "error",
			ping: func(_ context.Context) error {
				return errBadConn
			},
			expErr: errBadConn,
		},
		{
			name: "panic",
			ping: func(_ context.Context) error {
				panic("unexpected error")
			},
			expErr:    ErrConnectionLost,
			expErrMsg: "unexpected error",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err := Ping(ctx, tt.ping)
			require.ErrorIs(t, err, tt.expErr)
			if err != nil && tt.expErrMsg != "" {
				assert.Contains(t, err.Error(), tt.expErrMsg)
			}
		})
	}
}

type pingerMock func(ctx context.Context) error

func (m pingerMock) Ping(ctx context.Context) error {
	return m(ctx)
}
