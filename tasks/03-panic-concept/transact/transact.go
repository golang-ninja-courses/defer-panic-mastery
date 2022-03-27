package dbhelpers

type DB interface {
	Begin() (Tx, error)
}

type Tx interface {
	Exec(q string) error
	Commit() error
	Rollback() error
}

// Transact запускает функцию f в созданной для неё с помощью db транзакции.
// При успешном завершении функции транзакция фиксируется, иначе – откатывается.
func Transact(db DB, f func(tx Tx) error) error {
	// Реализуй меня.
	return nil
}
