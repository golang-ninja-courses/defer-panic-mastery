package errors

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttachTo_NoErrors(t *testing.T) {
	err := func() (err error) {
		defer AttachTo(&err, func() error { return nil })
		return nil
	}()
	require.NoError(t, err)
}

func TestAttachTo_NoDeferredError(t *testing.T) {
	err := func() (err error) {
		defer AttachTo(&err, func() error { return nil })
		return io.EOF
	}()
	require.Error(t, err)
	assert.ErrorIs(t, err, io.EOF)
	assert.EqualError(t, err, "EOF")
}

func TestAttachTo_NoFirstError(t *testing.T) {
	err := func() (err error) {
		defer AttachTo(&err, func() error { return io.EOF })
		return nil
	}()
	require.Error(t, err)
	assert.ErrorIs(t, err, io.EOF)
	assert.EqualError(t, err, "EOF")
}

func TestAttachTo_BothErrors(t *testing.T) {
	errBadConn := errors.New("bad connection")

	err := func() (err error) {
		defer AttachTo(&err, closerMock{errBadConn}.Close)
		return context.Canceled
	}()
	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
	assert.ErrorIs(t, err, errBadConn)
	assert.EqualError(t, err, "context canceled (and after bad connection)")
}

func TestAttachTo_Nil(t *testing.T) {
	assert.PanicsWithValue(t, "invalid usage of AttachTo", func() {
		_ = func() (err error) {
			defer AttachTo(nil, func() error { return io.EOF })
			return context.Canceled
		}()
	})
}

func TestAttachTo_NoFunc(t *testing.T) {
	assert.PanicsWithValue(t, "invalid usage of AttachTo", func() {
		_ = func() (err error) {
			defer AttachTo(&err, nil)
			return context.Canceled
		}()
	})
}

type closerMock struct {
	err error
}

func (m closerMock) Close() error {
	return m.err
}
