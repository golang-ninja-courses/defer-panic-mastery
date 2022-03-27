package errd

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	err := func() (err error) {
		defer Wrap(&err, "read from connection")

		if true {
			err = io.EOF
		}
		if true {
			err = io.ErrUnexpectedEOF
		}
		if false {
			return io.ErrClosedPipe
		}
		return context.Canceled
	}()

	require.Error(t, err)
	assert.EqualError(t, err, "read from connection: context canceled")
	assert.ErrorIs(t, err, context.Canceled)
}

func TestWrap_FormatStr(t *testing.T) {
	errExpiredToken := errors.New("expired token")

	err := func() (err error) {
		defer Wrap(&err, "user %q: authentication attempt %d", "@goinpractice", 3)
		err = errExpiredToken
		return
	}()

	require.Error(t, err)
	assert.EqualError(t, err, `user "@goinpractice": authentication attempt 3: expired token`)
	assert.ErrorIs(t, err, errExpiredToken)
}

func TestWrap_NoError(t *testing.T) {
	const machineID = 4242

	err := func() (err error) {
		defer Wrap(&err, "machine id %d", machineID)

		if true {
			err = io.EOF
		}
		if true {
			err = io.ErrUnexpectedEOF
		}
		return nil
	}()

	require.NoError(t, err)
}

func TestWrap_InvalidUsage(t *testing.T) {
	require.PanicsWithValue(t, "invalid usage of errd.Wrap", func() {
		_ = func() (err error) {
			defer Wrap(nil, "connect")
			return context.Canceled
		}()
	})
}
