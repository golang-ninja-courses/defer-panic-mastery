package fnhelpers

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	t.Run("ok no error 1", func(t *testing.T) {
		assert.NotPanics(t, func() {
			result := Must(func() (int, error) {
				return 42, nil
			})
			assert.Equal(t, 42, result)
		})
	})

	t.Run("ok no error 2", func(t *testing.T) {
		cmd := new(exec.Cmd)

		assert.NotPanics(t, func() {
			result := Must(func() (*exec.Cmd, error) {
				return cmd, nil
			})
			assert.Equal(t, cmd, result)
		})
	})

	t.Run("ok no error 3", func(t *testing.T) {
		assert.NotPanics(t, func() {
			result := Must(func() (*exec.Cmd, error) {
				return nil, nil //nolint:nilnil
			})
			assert.Nil(t, result)
		})
	})

	t.Run("panic because error 1", func(t *testing.T) {
		errExpected := errors.New("something went wrong")

		assert.PanicsWithValue(t, errExpected, func() {
			Must(func() (string, error) {
				return "", errExpected
			})
		})
	})

	t.Run("panic because error 2", func(t *testing.T) {
		errExpected := errors.New("something went wrong")

		assert.PanicsWithValue(t, errExpected, func() {
			Must(func() (string, error) {
				return "garbage", errExpected // Не делайте так в реальной жизни!
			})
		})
	})

	t.Run("panic while panic", func(t *testing.T) {
		expectedMessage := "something went really wrong"

		assert.PanicsWithValue(t, expectedMessage, func() {
			Must(func() (interface{}, error) {
				panic(expectedMessage)
			})
		})
	})
}
