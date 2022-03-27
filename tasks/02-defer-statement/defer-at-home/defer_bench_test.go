package deferr

import "testing"

/*
Для авторского решения:

$ go test -benchmem -bench .
Benchmark_withGoDefer-8          9457879               120.8 ns/op             8 B/op          1 allocs/op
Benchmark_withOwnDefer-8         7461952               155.8 ns/op           168 B/op          8 allocs/op
PASS
ok      tasks/02-defer-statement/defer-at-home    2.872s
*/

func Benchmark_withGoDefer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withGoDefer()
	}
}

func Benchmark_withOwnDefer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withOwnDefer()
	}
}

func withGoDefer() (result int) {
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)

	for i := 0; i < 3; i++ {
		defer foo(&result)
	}

	return
}

func withOwnDefer() (result int) {
	d := NewDefer()

	d.Defer(func() { foo(&result) })
	d.Defer(func() { foo(&result) })
	d.Defer(func() { foo(&result) })

	for i := 0; i < 3; i++ {
		d.Defer(func() { foo(&result) })
	}

	d.Execute()
	return
}

func foo(i *int) {
	*i++
}
