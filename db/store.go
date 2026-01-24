package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-movie/movie_resevation"
)

type Store interface {
	movie_resevation.Querier
	ExecTx(ctx context.Context, fn func(movie_resevation.Querier) error) error
}

type SQLStore struct {
	*movie_resevation.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: movie_resevation.New(db),
	}
}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(movie_resevation.Querier) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := movie_resevation.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
