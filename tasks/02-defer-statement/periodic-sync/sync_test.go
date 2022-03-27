package filehelpers

import (
	"context"
	"fmt"
	"log"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func ExampleSync() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second+100*time.Millisecond)
	defer cancel()

	vs := newVerboseSyncer(100 * time.Millisecond)
	err := Sync(ctx, vs, 200*time.Millisecond)
	if err != nil {
		panic(err)
	}

	// Output:
	// sync [elapsed from Sync() call – 200ms]
	// sync [elapsed from Sync() call – 400ms]
	// sync [elapsed from Sync() call – 600ms]
	// sync [elapsed from Sync() call – 800ms]
	// sync [elapsed from Sync() call – 1s]
}

func TestSync(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second+50*time.Millisecond)
	defer cancel()

	sm := new(syncerMock)
	errc := make(chan error, 1)
	go func() {
		errc <- Sync(ctx, sm, 100*time.Millisecond)
	}()

	select {
	case err := <-errc:
		require.NoError(t, err)
		assert.Equal(t, 10, sm.calls)

	case <-time.After(2 * time.Second):
		t.Fatal("Sync() was not cancelled by ctx")
	}
}

func TestSync_ManualCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sm := &syncerMock{
		actionCall: 5,
		action: func() error {
			cancel()
			return nil
		},
	}
	errc := make(chan error, 1)
	go func() {
		errc <- Sync(ctx, sm, 100*time.Millisecond)
	}()

	select {
	case err := <-errc:
		require.NoError(t, err)
		assert.Equal(t, 5, sm.calls)

	case <-time.After(2 * time.Second):
		t.Fatal("Sync() was not cancelled by ctx")
	}
}

func TestSync_Error(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second+50*time.Millisecond)
	defer cancel()

	sm := &syncerMock{
		actionCall: 5,
		action: func() error {
			return syscall.EBADF
		},
	}
	errc := make(chan error, 1)
	go func() {
		errc <- Sync(ctx, sm, 100*time.Millisecond)
	}()

	select {
	case err := <-errc:
		require.ErrorIs(t, err, syscall.EBADF)
		assert.Equal(t, 5, sm.calls)

	case <-time.After(2 * time.Second):
		t.Fatal("Sync() was not cancelled by ctx")
	}
}

type verboseSyncer struct {
	start    time.Time
	rounding time.Duration
}

func newVerboseSyncer(r time.Duration) *verboseSyncer {
	return &verboseSyncer{start: time.Now(), rounding: r}
}

func (v *verboseSyncer) Sync() error {
	fmt.Printf("sync [elapsed from Sync() call – %s]\n", time.Since(v.start).Round(v.rounding))
	return nil
}

type syncerMock struct {
	calls      int
	actionCall int
	action     func() error
}

func (s *syncerMock) Sync() error {
	s.calls++

	if s.actionCall != 0 && s.calls == s.actionCall {
		return s.action()
	}
	return nil
}
