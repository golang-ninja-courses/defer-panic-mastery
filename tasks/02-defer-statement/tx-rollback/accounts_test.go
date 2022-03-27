package accounts

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	var (
		errNoSuchTable        = errors.New("no such table")
		errTableAlreadyExists = errors.New("table already exists")
		errDuplicatedKey      = errors.New("duplicate key value violates unique constraint")
		errBadConnection      = errors.New("driver: bad connection")
		errAbortedTx          = errors.New("current transaction is aborted")
	)

	cases := []struct {
		name                                                                                        string
		beginErr, dropErr, createErr, fillErr, commitErr, rollbackErr                               error
		expBeginCalls, expDropCalls, expCreateCalls, expFillCalls, expCommitCalls, expRollbackCalls int
		expCreateErr                                                                                error
	}{
		{
			name:           "positive scenario",
			expBeginCalls:  1,
			expDropCalls:   1,
			expCreateCalls: 1,
			expFillCalls:   1,
			expCommitCalls: 1,
			expCreateErr:   nil,
		},
		{
			name:          "begin error",
			beginErr:      errBadConnection,
			expBeginCalls: 1,
			expCreateErr:  errBadConnection,
		},
		{
			name:             "drop error",
			dropErr:          errNoSuchTable,
			expBeginCalls:    1,
			expDropCalls:     1,
			expRollbackCalls: 1,
			expCreateErr:     errNoSuchTable,
		},
		{
			name:             "rollback err after drop error",
			dropErr:          errNoSuchTable,
			rollbackErr:      errBadConnection,
			expBeginCalls:    1,
			expDropCalls:     1,
			expRollbackCalls: 1,
			expCreateErr:     errBadConnection,
		},
		{
			name:             "create error",
			createErr:        errTableAlreadyExists,
			expBeginCalls:    1,
			expDropCalls:     1,
			expCreateCalls:   1,
			expRollbackCalls: 1,
			expCreateErr:     errTableAlreadyExists,
		},
		{
			name:             "rollback err after create error",
			createErr:        errTableAlreadyExists,
			rollbackErr:      errBadConnection,
			expBeginCalls:    1,
			expDropCalls:     1,
			expCreateCalls:   1,
			expRollbackCalls: 1,
			expCreateErr:     errBadConnection,
		},
		{
			name:             "fill error",
			fillErr:          errDuplicatedKey,
			expBeginCalls:    1,
			expDropCalls:     1,
			expCreateCalls:   1,
			expFillCalls:     1,
			expRollbackCalls: 1,
			expCreateErr:     errDuplicatedKey,
		},
		{
			name:             "rollback err after fill error",
			fillErr:          errDuplicatedKey,
			rollbackErr:      errBadConnection,
			expBeginCalls:    1,
			expDropCalls:     1,
			expCreateCalls:   1,
			expFillCalls:     1,
			expRollbackCalls: 1,
			expCreateErr:     errBadConnection,
		},
		{
			name:           "commit error",
			commitErr:      errAbortedTx,
			expBeginCalls:  1,
			expDropCalls:   1,
			expCreateCalls: 1,
			expFillCalls:   1,
			expCommitCalls: 1,
			expCreateErr:   errAbortedTx,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tx := &txMock{
				DropErr:     tt.dropErr,
				CreateErr:   tt.createErr,
				FillErr:     tt.fillErr,
				CommitErr:   tt.commitErr,
				RollbackErr: tt.rollbackErr,
			}
			db := &dbMock{
				BeginErr: tt.beginErr,
				Tx:       tx,
			}
			err := Create(db)
			require.ErrorIs(t, err, tt.expCreateErr)

			assert.Equal(t, tt.expBeginCalls, db.BeginCalls)
			assert.Equal(t, tt.expDropCalls, tx.DropCalls)
			assert.Equal(t, tt.expCreateCalls, tx.CreateCalls)
			assert.Equal(t, tt.expFillCalls, tx.FillCalls)
			assert.Equal(t, tt.expCommitCalls, tx.CommitCalls)
			assert.Equal(t, tt.expRollbackCalls, tx.RollbackCalls)
		})
	}
}

var (
	_ DB = new(dbMock)
	_ Tx = new(txMock)
)

type dbMock struct {
	BeginErr   error
	BeginCalls int
	Tx         *txMock
}

func (db *dbMock) Begin() (Tx, error) {
	db.BeginCalls++

	if db.BeginErr != nil {
		return nil, db.BeginErr
	}
	return db.Tx, nil
}

type txMock struct {
	DropErr     error
	CreateErr   error
	FillErr     error
	CommitErr   error
	RollbackErr error

	DropCalls     int
	CreateCalls   int
	FillCalls     int
	CommitCalls   int
	RollbackCalls int
}

func (t *txMock) Exec(q string) error {
	switch q {
	case queryDropAccounts:
		t.DropCalls++
		return t.DropErr

	case queryCreateAccounts:
		t.CreateCalls++
		return t.CreateErr

	case queryFillAccounts:
		t.FillCalls++
		return t.FillErr

	default:
		return fmt.Errorf("unexpected exec: %q", q)
	}
}

func (t *txMock) Commit() error {
	t.CommitCalls++
	return t.CommitErr
}

func (t *txMock) Rollback() error {
	t.RollbackCalls++
	return t.RollbackErr
}
