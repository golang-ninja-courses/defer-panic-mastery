package coolassertlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsserts(t *testing.T) {
	cases := []struct {
		name        string
		fn          func()
		expectPanic bool
	}{
		{
			name:        "no panic",
			fn:          func() {},
			expectPanic: false,
		},
		{
			name: "panic nil",
			fn: func() {
				panic(nil)
			},
			expectPanic: true,
		},
		{
			name: "panic not nil",
			fn: func() {
				panic("boom!")
			},
			expectPanic: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("AssertPanics", func(t *testing.T) {
				tMock := new(mockT)
				wasPanic := AssertPanics(tMock, tt.fn)
				assert.Equal(t, tt.expectPanic, wasPanic)
				assert.Equal(t, !tt.expectPanic, tMock.failed, "check t.Errorf() call")
				assert.True(t, tMock.markedAsHelper, "no t.Helper() call")
			})

			t.Run("AssertNotPanics", func(t *testing.T) {
				tMock := new(mockT)
				wasNotPanic := AssertNotPanics(tMock, tt.fn)
				assert.Equal(t, !tt.expectPanic, wasNotPanic)
				assert.Equal(t, tt.expectPanic, tMock.failed, "check t.Errorf() call")
				assert.True(t, tMock.markedAsHelper, "no t.Helper() call")
			})
		})
	}
}

type mockT struct {
	failed         bool
	markedAsHelper bool
}

func (t *mockT) Helper() {
	t.markedAsHelper = true
}

func (t *mockT) Errorf(format string, args ...interface{}) {
	t.failed = true
}
