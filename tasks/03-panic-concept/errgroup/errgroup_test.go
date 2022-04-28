package errgroup

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestGroup(t *testing.T) {
	t.Run("verify error handling", func(t *testing.T) {
		defer goleak.VerifyNone(t)

		err1 := errors.New("err 1")
		err2 := errors.New("err 2")

		cases := []struct {
			errs        []error
			errExpected error
		}{
			{errs: []error{}, errExpected: nil},
			{errs: []error{nil}, errExpected: nil},
			{errs: []error{err1}, errExpected: err1},
			{errs: []error{err1, nil}, errExpected: err1},
			{errs: []error{err1, nil, err2}, errExpected: err1},
			{errs: []error{err2, err1, nil}, errExpected: err2},
			{errs: []error{nil, nil, nil, err2, err1}, errExpected: err2},
		}

		for _, tt := range cases {
			var g Group

			for i, err := range tt.errs {
				i, err := i, err

				g.Go(func() error {
					time.Sleep(time.Duration(i) * 50 * time.Millisecond)
					return err
				})
			}

			gErr := g.Wait()
			assert.Equal(t, tt.errExpected, gErr) // Not ErrorIs!
		}
	})

	// verify panic handling
	{
		const expectedPrefix = "recovered panic: "

		type User struct {
			ID string
		}

		errExpected := errors.New("expected error")

		cases := []struct {
			name                      string
			panicValue                any
			errExpected               error
			expectedInitialErrMessage string
		}{
			{
				name:        "panicking with error",
				panicValue:  errExpected,
				errExpected: errExpected,
			},
			{
				name:                      "panicking with string",
				panicValue:                "expected message",
				expectedInitialErrMessage: "expected message",
			},
			{
				name:                      "panicking with int",
				panicValue:                12345,
				expectedInitialErrMessage: "12345",
			},
			{
				name:                      "panicking with struct",
				panicValue:                User{ID: "XXX"},
				expectedInitialErrMessage: `errgroup.User{ID:"XXX"}`,
			},
		}

		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				var g Group
				g.Go(func() error { panic(tt.panicValue) })

				err := g.Wait()
				require.Error(t, err)
				require.True(t, strings.HasPrefix(err.Error(), expectedPrefix))

				if tt.errExpected != nil {
					assert.ErrorIs(t, err, tt.errExpected)
				}

				if tt.expectedInitialErrMessage != "" {
					assert.Equal(t, err.Error()[len(expectedPrefix):], tt.expectedInitialErrMessage)
				}
			})
		}
	}

	t.Run("verify panic after error handling", func(t *testing.T) {
		errInvalidIntegrity := errors.New("invalid integrity")

		var g Group

		g.Go(func() error {
			return errInvalidIntegrity
		})
		g.Go(func() error {
			time.Sleep(50 * time.Millisecond)
			return nil
		})
		g.Go(func() error {
			time.Sleep(100 * time.Millisecond)
			panic("any message")
		})

		err := g.Wait()
		assert.ErrorIs(t, err, errInvalidIntegrity)
	})

	t.Run("verify error after panic handling", func(t *testing.T) {
		errInvalidIntegrity := errors.New("invalid integrity")

		var g Group

		g.Go(func() error {
			panic("any message")
		})
		g.Go(func() error {
			time.Sleep(50 * time.Millisecond)
			return nil
		})
		g.Go(func() error {
			time.Sleep(100 * time.Millisecond)
			return errInvalidIntegrity
		})

		err := g.Wait()
		assert.EqualError(t, err, "recovered panic: any message")
	})

	t.Run("cancelling context when error occurs", func(t *testing.T) {
		defer goleak.VerifyNone(t)

		errInvalidIntegrity := errors.New("invalid integrity")

		var badGoReturned bool
		g, ctx := WithContext(context.Background())

		g.Go(func() error {
			select {
			case <-time.After(time.Second):
				t.Fatal("no expected ctx cancellation")

			case <-ctx.Done():
				assert.True(t, badGoReturned)
			}
			return nil
		})

		g.Go(func() error {
			defer func() { badGoReturned = true }()
			return errInvalidIntegrity
		})

		err := g.Wait()
		assert.ErrorIs(t, err, errInvalidIntegrity)
	})

	t.Run("cancelling context when panic occurs", func(t *testing.T) {
		defer goleak.VerifyNone(t)

		var badGoReturned bool
		g, ctx := WithContext(context.Background())

		g.Go(func() error {
			select {
			case <-time.After(time.Second):
				t.Fatal("no expected ctx cancellation")

			case <-ctx.Done():
				assert.True(t, badGoReturned)
				assert.ErrorIs(t, ctx.Err(), context.Canceled)
			}
			return nil
		})

		g.Go(func() error {
			defer func() { badGoReturned = true }()
			panic("any message")
		})

		err := g.Wait()
		assert.EqualError(t, err, "recovered panic: any message")
	})
}
