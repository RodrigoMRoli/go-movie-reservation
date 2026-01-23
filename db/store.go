package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-movie-reservation/movie_resevation"
)

type Store interface {
	movie_resevation.Querier // Interface gerada pelo sqlc que tem todos os métodos (CreateMovie, etc)
	CreateMovieWithGenre(ctx context.Context, arg movie_resevation.CreateMovieParams) error
}

type SQLStore struct {
	*movie_resevation.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: movie_resevation.New(db), // Função gerada pelo sqlc
	}
}

// ExecTx é uma função auxiliar para executar qualquer coisa dentro de uma transação
func (store *SQLStore) execTx(ctx context.Context, fn func(*movie_resevation.Queries) error) error {
	// 1. Inicia a transação no banco
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 2. Cria uma nova instância de Queries que usa ESSA transação
	q := movie_resevation.New(tx)

	// 3. Roda a função que passamos (nossa lógica de negócio)
	err = fn(q)

	// 4. Decide se faz Commit ou Rollback
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr) // Erro duplo (na lógica e no rollback)
		}
		return err
	}

	return tx.Commit()
}
