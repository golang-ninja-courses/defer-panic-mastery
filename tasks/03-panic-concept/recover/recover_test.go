package panichelpers

import (
	"runtime"
	"testing"
	"text/template"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecover(t *testing.T) {
	cases := []struct {
		name   string
		newArg func() any
	}{
		{
			name:   "bool",
			newArg: func() any { return false },
		},
		{
			name:   "int",
			newArg: func() any { return 42 },
		},
		{
			name:   "int8",
			newArg: func() any { return int8(42) },
		},
		{
			name:   "int16",
			newArg: func() any { return int16(42) },
		},
		{
			name:   "int32",
			newArg: func() any { return int32(42) },
		},
		{
			name:   "int64",
			newArg: func() any { return int64(42) },
		},
		{
			name:   "uint",
			newArg: func() any { return uint(24) },
		},
		{
			name:   "uint8",
			newArg: func() any { return uint8(24) },
		},
		{
			name:   "uint16",
			newArg: func() any { return uint16(24) },
		},
		{
			name:   "uint32",
			newArg: func() any { return uint32(24) },
		},
		{
			name:   "uint64",
			newArg: func() any { return uint64(24) },
		},
		{
			name:   "float32",
			newArg: func() any { return float32(42.) },
		},
		{
			name:   "float64",
			newArg: func() any { return 42. },
		},
		{
			name:   "complex64",
			newArg: func() any { return complex(float32(1.), float32(10.)) },
		},
		{
			name:   "complex128",
			newArg: func() any { return complex(10., 100.) },
		},
		{
			name:   "array",
			newArg: func() any { return [...]int{1, 2, 3, 4, 5} },
		},
		{
			name:   "chan",
			newArg: func() any { return make(chan int, 100) },
		},
		{
			name:   "interface",
			newArg: func() any { return error(new(runtime.TypeAssertionError)) },
		},
		{
			name: "ptr",
			newArg: func() any {
				i := 42
				return &i
			},
		},
		{
			name:   "string",
			newArg: func() any { return "don't panic!" },
		},
		{
			name:   "struct",
			newArg: func() any { return template.ExecError{} },
		},
		{
			name: "unsafe ptr",
			newArg: func() any {
				i := 42
				return unsafe.Pointer(&i)
			},
		},
		// func, map & slice are not comparable.
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				var fnCalled bool
				arg := tt.newArg()
				p := Recover(func() {
					fnCalled = true
					panic(arg)
				})
				require.NotNil(t, p)
				assert.True(t, fnCalled)
				assert.Equal(t, arg, p)
			})
		})
	}
}

func TestRecover_PanicWithNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var fnCalled bool
		p := Recover(func() {
			fnCalled = true
			panic(nil)
		})
		require.Nil(t, p)
		assert.True(t, fnCalled)
	})
}

func TestRecover_NoPanic(t *testing.T) {
	var fnCalled bool
	p := Recover(func() {
		fnCalled = true
	})
	require.Nil(t, p)
	assert.True(t, fnCalled)
}
