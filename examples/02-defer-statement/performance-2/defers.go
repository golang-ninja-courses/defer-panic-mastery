package main

//go:noinline
func withoutDefer() (result int) {
	foo(&result)
	return
}

//go:noinline
func withDefer() (result int) {
	defer foo(&result)
	return
}

//go:noinline
func withDefers() (result int) {
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	return
}

//go:noinline
func withNotOpenDefer() (result int) {
	for i := 0; i < 1; i++ {
		defer foo(&result)
	}
	return
}

//go:noinline
func with8Defers() (result int) {
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	return
}

//go:noinline
func with9Defers() (result int) {
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	defer foo(&result)
	return
}

func foo(i *int) {
	*i++
}
