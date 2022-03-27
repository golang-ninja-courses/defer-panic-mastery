package db

import "fmt"

func Operation(d DB) {
	d.RunInTransaction()
	execQuery(d)
}

func execQuery(_ DB) {
	fmt.Println("EXEC QUERY")
}

type DB struct{}

// RunInTransaction создаёт транзакцию и коммитит её,
// когда завершается функция, в которой произошёл вызов Trans().
func (d DB) RunInTransaction() {
	d.Begin()
	defer d.Commit()
}

func (d DB) Begin() {
	fmt.Println("BEGIN TRANSACTION")
}

func (d DB) Commit() {
	fmt.Println("COMMIT TRANSACTION")
}
