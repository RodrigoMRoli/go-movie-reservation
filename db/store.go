package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-movie-reservation/movie_resevation"
	// Importe o pacote gerado pelo sqlc
)

// Store define todas as funções de banco de dados disponíveis
// Ela "herda" a interface Querier (gerada pelo sqlc) e adiciona ExecTx
type Store interface {
	movie_resevation.Querier
	ExecTx(ctx context.Context, fn func(movie_resevation.Querier) error) error
}

// SQLStore é a implementação real
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

// ExecTx executa uma função dentro de uma transação do banco
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
