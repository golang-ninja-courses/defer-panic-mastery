package comparators

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceEqual(t *testing.T) {
	cases := []struct {
		name     string
		lhs, rhs any
		expected bool
	}{
		{
			name:     "nil notnil",
			lhs:      nil,
			rhs:      new(int),
			expected: false,
		},
		{
			name:     "notnil nil",
			lhs:      new(int),
			rhs:      nil,
			expected: false,
		},
		{
			name:     "different errors",
			lhs:      io.EOF,
			rhs:      io.ErrUnexpectedEOF,
			expected: false,
		},
		{
			name:     "different dynamic types",
			lhs:      "hello",
			rhs:      42,
			expected: false,
		},
		{
			name:     "incomparable dynamic types",
			lhs:      http.HandlerFunc(nil),
			rhs:      http.HandlerFunc(nil),
			expected: false,
		},
		{
			name:     "nil nil",
			lhs:      nil,
			rhs:      nil,
			expected: true,
		},
		{
			name:     "identical errors",
			lhs:      io.EOF,
			rhs:      io.EOF,
			expected: true,
		},
		{
			name:     "identical strings",
			lhs:      "gogo",
			rhs:      "gogo",
			expected: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				b := InterfaceEqual(tt.lhs, tt.rhs)
				assert.Equal(t, tt.expected, b)
			})
		})
	}
}
