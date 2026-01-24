package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/auth"
)

type Store interface {
	auth.Querier
	ExecTx(ctx context.Context, fn func(auth.Querier) error) error
}

type SQLStore struct {
	*auth.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: auth.New(db),
	}
}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(auth.Querier) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := auth.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
