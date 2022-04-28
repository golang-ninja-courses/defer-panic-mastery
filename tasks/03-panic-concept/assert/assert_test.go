package assert

import (
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssert_Enabled(t *testing.T) {
	cmd := exec.Command("go", "run", filepath.Join("testdata", "main.go"))
	out, err := cmd.CombinedOutput()

	var exitErr *exec.ExitError
	require.ErrorAs(t, err, &exitErr)
	require.False(t, exitErr.Success())
	require.Contains(t, string(out), "panic: v must be initialized")
}

func TestAssert_Disabled(t *testing.T) {
	cmd := exec.Command("go", "run", "-tags", "NDEBUG", filepath.Join("testdata", "main.go"))
	out, err := cmd.CombinedOutput()
	require.NoError(t, err)
	require.Contains(t, string(out), "I'm OK")
}

func TestAssert(t *testing.T) {
	cases := []struct {
		cond          bool
		msg           string
		args          []any
		expectedPanic string
	}{
		{
			cond:          true,
			msg:           "identifier already declared or resolved",
			expectedPanic: "",
		},
		{
			cond:          true,
			msg:           "identifiers %q, %d, %#v already declared or resolved",
			args:          []any{"133a656", 4242, struct{ string }{"64ebed8"}},
			expectedPanic: "",
		},
		{
			cond:          false,
			msg:           "identifier already declared or resolved",
			expectedPanic: "identifier already declared or resolved",
		},
		{
			cond:          false,
			msg:           "identifiers %q, %d, %#v already declared or resolved",
			args:          []any{"133a656", 4242, struct{ string }{"64ebed8"}},
			expectedPanic: `identifiers "133a656", 4242, struct { string }{string:"64ebed8"} already declared or resolved`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			if tt.expectedPanic != "" {
				assert.PanicsWithValue(t, tt.expectedPanic, func() {
					Assert(tt.cond, tt.msg, tt.args...)
				})
			} else {
				assert.NotPanics(t, func() {
					Assert(tt.cond, tt.msg, tt.args...)
				})
			}
		})
	}
}
