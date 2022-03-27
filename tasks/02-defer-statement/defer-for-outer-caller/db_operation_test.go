package db

func ExampleOperation() {
	var d DB
	Operation(d)

	// Output:
	// BEGIN TRANSACTION
	// EXEC QUERY
	// COMMIT TRANSACTION
}
