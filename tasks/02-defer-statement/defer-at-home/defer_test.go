package deferr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleDefer() {
	d := NewDefer()
	for _, c := range "hello" {
		c := c
		d.Defer(func() { fmt.Print(string(c)) })
	}
	d.Execute()

	// Output:
	// olleh
}

func TestDefer(t *testing.T) {
	var results []int

	d := NewDefer()
	for i := 1; i <= 3; i++ {
		i := i
		d.Defer(func() {
			results = append(results, i)
		})
	}

	f := func() {
		results = append(results, 4)
	}
	d.Defer(f)

	require.Empty(t, results)
	results = append(results, 0)

	d.Execute()
	require.Equal(t, []int{0, 4, 3, 2, 1}, results)
}

func TestDefer_NoFuncs(t *testing.T) {
	NewDefer().Execute()
}

func TestDefer_NoGlobalStack(t *testing.T) {
	d1, d2 := NewDefer(), NewDefer()

	var results1 []int
	var results2 []int

	for i := 1; i <= 5; i++ {
		i := i
		d1.Defer(func() { results1 = append(results1, i) })
		d2.Defer(func() { results2 = append(results2, i+10) })
	}

	require.Empty(t, results1)
	require.Empty(t, results2)

	d1.Execute()
	require.Equal(t, []int{5, 4, 3, 2, 1}, results1)
	require.Empty(t, results2)

	d2.Execute()
	require.Equal(t, []int{5, 4, 3, 2, 1}, results1)
	require.Equal(t, []int{15, 14, 13, 12, 11}, results2)
}
