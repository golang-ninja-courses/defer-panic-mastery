package fnhelpers

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatHappened(t *testing.T) {
	cases := []struct {
		name     string
		fn       func()
		expected ExitReason
	}{
		{
			name:     "fn does nothing",
			fn:       func() {},
			expected: ExitReasonRegularReturn,
		},
		{
			name: "fn calls panic",
			fn: func() {
				panic("boom!")
			},
			expected: ExitReasonPanic,
		},
		{
			name: "fn calls panic(nil)",
			fn: func() {
				panic(nil)
			},
			expected: ExitReasonPanic,
		},
		{
			name: "fn calls runtime.Goexit",
			fn: func() {
				runtime.Goexit()
			},
			expected: ExitReasonGoexit,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			reason := WhatHappened(tt.fn)
			assert.Equal(t, tt.expected, reason)
		})
	}
}
