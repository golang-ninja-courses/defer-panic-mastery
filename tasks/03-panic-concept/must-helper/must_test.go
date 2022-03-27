package fnhelpers

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	t.Run("ok nil result nil error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			result := Must(func() (interface{}, error) {
				return nil, nil //nolint:nilnil
			})
			assert.Nil(t, result)
		})
	})

	t.Run("ok notnil result nil error", func(t *testing.T) {
		cmd := new(exec.Cmd)

		assert.NotPanics(t, func() {
			result := Must(func() (interface{}, error) {
				return cmd, nil
			})
			assert.Equal(t, cmd, result)
		})
	})

	t.Run("panic nil result notnil error", func(t *testing.T) {
		errExpected := errors.New("something went wrong")

		assert.PanicsWithValue(t, errExpected, func() {
			Must(func() (interface{}, error) {
				return nil, errExpected
			})
		})
	})

	t.Run("panic notnil result notnil error", func(t *testing.T) {
		cmd := new(exec.Cmd)
		errExpected := errors.New("something went wrong")

		assert.PanicsWithValue(t, errExpected, func() {
			Must(func() (interface{}, error) {
				return cmd, errExpected // Не делайте так в реальной жизни!
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
