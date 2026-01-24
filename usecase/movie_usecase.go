package usecase

import (
	"context"
	"go-movie-reservation/model"
	"go-movie-reservation/repository"

	"github.com/google/uuid"
)

type MovieUseCase struct {
	repository repository.MovieRepository
}

func NewMovieUseCase(repository repository.MovieRepository) MovieUseCase {
	return MovieUseCase{
		repository: repository,
	}
}

func (mu *MovieUseCase) GetMovies(ctx context.Context) ([]model.MovieWithGenre, error) {
	return mu.repository.GetMovies(ctx)
}

func (mu *MovieUseCase) GetMovie(ctx context.Context, id uuid.UUID) (model.MovieWithGenre, error) {
	return mu.repository.GetMovie(ctx, id)
}

func (mu *MovieUseCase) CreateMovie(ctx context.Context, params model.CreateMovieInput) (model.MovieWithGenre, error) {
	return mu.repository.CreateMovie(ctx, params)
}

func (mu *MovieUseCase) UpdateMovie(ctx context.Context, id uuid.UUID, params model.UpdateMovieInput) (model.MovieWithGenre, error) {
	return mu.repository.UpdateMovie(ctx, id, params)
}

func (mu *MovieUseCase) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	return mu.repository.DeleteMovie(ctx, id)
}
