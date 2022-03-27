package main

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

//go:noinline
func with9DefersHack() (result int) {
	func() {
		defer foo(&result)
		defer foo(&result)
		defer foo(&result)
		defer foo(&result)
		defer foo(&result)
		defer foo(&result)
		defer foo(&result)
		defer foo(&result)
	}()
	defer foo(&result)
	return
}

func foo(i *int) {
	*i++
}
