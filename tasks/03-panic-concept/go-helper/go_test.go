package errhelpers

import (
	"errors"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestGo(t *testing.T) {
	errInvalidIntegrity := errors.New("invalid integrity")

	cases := []struct {
		name           string
		fn             func() error
		errExpected    error
		errMsgExpected string
	}{
		{
			name:        "no error",
			fn:          func() error { return nil },
			errExpected: nil,
		},
		{
			name:        "with error",
			fn:          func() error { return errInvalidIntegrity },
			errExpected: errInvalidIntegrity,
		},
		{
			name:           "with panic",
			fn:             func() error { panic("boom!") },
			errExpected:    ErrPanicOccurred,
			errMsgExpected: "boom!",
		},
		{
			name: "with panic after error",
			fn: func() (err error) {
				err = errInvalidIntegrity //nolint:wastedassign
				panic("boom!")
			},
			errExpected: ErrPanicOccurred,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			errc := Go(tt.fn)
			select {
			case err := <-errc:
				assert.ErrorIs(t, err, tt.errExpected)
				if tt.errMsgExpected != "" {
					require.Error(t, err)
					assert.Contains(t, err.Error(), tt.errMsgExpected)
				}

			case <-time.After(time.Second):
				t.Fatal("unexpected blocking")
			}
		})
	}
}

func TestGo_NoResultReading(t *testing.T) {
	_ = Go(func() error {
		return io.EOF
	})
}
