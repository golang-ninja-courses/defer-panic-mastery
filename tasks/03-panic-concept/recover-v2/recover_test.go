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
		newArg func() interface{}
	}{
		{
			name:   "bool",
			newArg: func() interface{} { return false },
		},
		{
			name:   "int",
			newArg: func() interface{} { return 42 },
		},
		{
			name:   "int8",
			newArg: func() interface{} { return int8(42) },
		},
		{
			name:   "int16",
			newArg: func() interface{} { return int16(42) },
		},
		{
			name:   "int32",
			newArg: func() interface{} { return int32(42) },
		},
		{
			name:   "int64",
			newArg: func() interface{} { return int64(42) },
		},
		{
			name:   "uint",
			newArg: func() interface{} { return uint(24) },
		},
		{
			name:   "uint8",
			newArg: func() interface{} { return uint8(24) },
		},
		{
			name:   "uint16",
			newArg: func() interface{} { return uint16(24) },
		},
		{
			name:   "uint32",
			newArg: func() interface{} { return uint32(24) },
		},
		{
			name:   "uint64",
			newArg: func() interface{} { return uint64(24) },
		},
		{
			name:   "float32",
			newArg: func() interface{} { return float32(42.) },
		},
		{
			name:   "float64",
			newArg: func() interface{} { return 42. },
		},
		{
			name:   "complex64",
			newArg: func() interface{} { return complex(float32(1.), float32(10.)) },
		},
		{
			name:   "complex128",
			newArg: func() interface{} { return complex(10., 100.) },
		},
		{
			name:   "array",
			newArg: func() interface{} { return [...]int{1, 2, 3, 4, 5} },
		},
		{
			name:   "chan",
			newArg: func() interface{} { return make(chan int, 100) },
		},
		{
			name:   "interface",
			newArg: func() interface{} { return error(new(runtime.TypeAssertionError)) },
		},
		{
			name: "ptr",
			newArg: func() interface{} {
				i := 42
				return &i
			},
		},
		{
			name:   "string",
			newArg: func() interface{} { return "don't panic!" },
		},
		{
			name:   "struct",
			newArg: func() interface{} { return template.ExecError{} },
		},
		{
			name: "unsafe ptr",
			newArg: func() interface{} {
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
				p, wasPanic := Recover(func() {
					fnCalled = true
					panic(arg)
				})
				require.NotNil(t, p)
				assert.True(t, fnCalled)
				assert.Equal(t, arg, p)
				assert.True(t, wasPanic)
			})
		})
	}
}

func TestRecover_PanicWithNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var fnCalled bool
		p, wasPanic := Recover(func() {
			fnCalled = true
			panic(nil)
		})
		require.Nil(t, p)
		assert.True(t, fnCalled)
		assert.True(t, wasPanic)
	})
}

func TestRecover_NoPanic(t *testing.T) {
	var fnCalled bool
	p, wasPanic := Recover(func() {
		fnCalled = true
	})
	require.Nil(t, p)
	assert.True(t, fnCalled)
	assert.False(t, wasPanic)
}
