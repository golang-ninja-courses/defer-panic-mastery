package dbhelpers

import (
	"errors"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransact(t *testing.T) {
	var (
		errBadConnection = errors.New("driver: bad connection")
		errDuplicatedKey = errors.New("duplicate key value violates unique constraint")
		errAbortedTx     = errors.New("current transaction is aborted")
	)

	cases := []struct {
		name                                                        string
		wantPanic                                                   bool
		beginErr, fnErr, commitErr, rollbackErr                     error
		expBeginCalls, expFnCalls, expCommitCalls, expRollbackCalls int
		expectedTransactErr                                         error
	}{
		{
			name:                "begin error",
			beginErr:            errBadConnection,
			expBeginCalls:       1,
			expectedTransactErr: errBadConnection,
		},
		{
			name:                "fn err",
			fnErr:               errDuplicatedKey,
			expBeginCalls:       1,
			expFnCalls:          1,
			expRollbackCalls:    1,
			expectedTransactErr: errDuplicatedKey,
		},
		{
			name:                "fn err rollback err",
			fnErr:               errDuplicatedKey,
			rollbackErr:         errAbortedTx,
			expBeginCalls:       1,
			expFnCalls:          1,
			expRollbackCalls:    1,
			expectedTransactErr: errDuplicatedKey,
		},
		{
			name:             "fn panic",
			wantPanic:        true,
			expBeginCalls:    1,
			expFnCalls:       1,
			expRollbackCalls: 1,
		},
		{
			name:             "fn panic rollback err",
			wantPanic:        true,
			rollbackErr:      errAbortedTx,
			expBeginCalls:    1,
			expFnCalls:       1,
			expRollbackCalls: 1,
		},
		{
			name:                "fn success commit err",
			commitErr:           errAbortedTx,
			expBeginCalls:       1,
			expFnCalls:          1,
			expCommitCalls:      1,
			expectedTransactErr: errAbortedTx,
		},
		{
			name:           "fn success commit success",
			expBeginCalls:  1,
			expFnCalls:     1,
			expCommitCalls: 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tx := &txMock{
				CommitErr:   tt.commitErr,
				RollbackErr: tt.rollbackErr,
			}
			db := &dbMock{
				BeginErr: tt.beginErr,
				Tx:       tx,
			}

			var fnCalls int
			panicValue := new(runtime.TypeAssertionError)
			fn := func(_ Tx) error {
				fnCalls++

				if tt.wantPanic {
					panic(panicValue)
				}
				return tt.fnErr
			}

			if tt.wantPanic {
				require.PanicsWithValue(t, panicValue, func() {
					_ = Transact(db, fn)
				})
			} else {
				require.NotPanics(t, func() {
					err := Transact(db, fn)
					require.ErrorIs(t, err, tt.expectedTransactErr)
				})
			}

			assert.Equal(t, tt.expBeginCalls, db.BeginCalls)
			assert.Equal(t, tt.expFnCalls, fnCalls)
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
	CommitErr   error
	CommitCalls int

	RollbackErr   error
	RollbackCalls int
}

func (t *txMock) Exec(_ string) error {
	// Unused method, just for example.
	return nil
}

func (t *txMock) Commit() error {
	t.CommitCalls++
	return t.CommitErr
}

func (t *txMock) Rollback() error {
	t.RollbackCalls++
	return t.RollbackErr
}
