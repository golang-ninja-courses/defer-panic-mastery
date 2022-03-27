package logger

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleLogError() {
	LogError(sql.ErrConnDone, os.Stdout)

	// Output:
	// error occurred: sql: connection is already closed
}

func TestLogError(t *testing.T) {
	cases := []struct {
		name           string
		err            error
		expectedOutput string
	}{
		{
			name:           "valid sentinel error",
			err:            io.EOF,
			expectedOutput: "error occurred: EOF",
		},
		{
			name:           "valid inplace created error",
			err:            errors.New("bad connection"),
			expectedOutput: "error occurred: bad connection",
		},
		{
			name:           "nil error",
			err:            nil,
			expectedOutput: "error occurred: <nil>",
		},
		{
			name:           "broken error",
			err:            brokenError{},
			expectedOutput: "error occurred: (logger.brokenError).Error() panic: runtime error: invalid memory address or nil pointer dereference",
		},
		{
			name:           "explicitly panicking error",
			err:            new(panickingError),
			expectedOutput: "error occurred: (*logger.panickingError).Error() panic: no such process",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				var b bytes.Buffer
				LogError(tt.err, &b)
				assert.Equal(t, tt.expectedOutput, b.String())
			})
		})
	}
}

type brokenError struct {
	error
}

type panickingError struct{}

func (err *panickingError) Error() string {
	return fmt.Sprintf("process %d is died", getPID())
}

func getPID() int {
	panic(syscall.ESRCH)
}
