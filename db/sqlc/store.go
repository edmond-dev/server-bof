package db

import (
	"context"
	"database/sql"
	"fmt"
)

//Store provides all functions to execute queries db and transactions
type Store struct {
	*Queries
	db *sql.DB
}

//NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

//ExecTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf(`tx err: %v, rb err: %v`, err, rberr)
		}
		return err
	}
	return tx.Commit()
}
