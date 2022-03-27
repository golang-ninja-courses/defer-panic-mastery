package txhelpers

import (
	"database/sql"
	"log"
)

func RunInTransaction(tx *sql.Tx, fn func(*sql.Tx) error) error {
	if err := fn(tx); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			log.Println("tx.Rollback failed:", rErr)
			// err = fmt.Errorf("%w and after rollback error: %v", err, rErr)
		}
		return err
	}
	return tx.Commit()
}
