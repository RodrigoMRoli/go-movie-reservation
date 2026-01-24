package usecase

import (
	"context"
	"go-movie-reservation/model"
	"go-movie-reservation/service"

	"github.com/google/uuid"
)

type MovieUseCase struct {
	service service.MovieService
}

func NewMovieUseCase(service service.MovieService) MovieUseCase {
	return MovieUseCase{
		service: service,
	}
}

func (mu *MovieUseCase) GetMovies(ctx context.Context) ([]model.MovieWithGenre, error) {
	return mu.service.GetMovies(ctx)
}

func (mu *MovieUseCase) GetMovie(ctx context.Context, id uuid.UUID) (model.MovieWithGenre, error) {
	return mu.service.GetMovie(ctx, id)
}

func (mu *MovieUseCase) CreateMovie(ctx context.Context, params model.CreateMovieInput) (model.MovieWithGenre, error) {
	return mu.service.CreateMovie(ctx, params)
}

func (mu *MovieUseCase) UpdateMovie(ctx context.Context, id uuid.UUID, params model.UpdateMovieInput) (model.MovieWithGenre, error) {
	return mu.service.UpdateMovie(ctx, id, params)
}

func (mu *MovieUseCase) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	return mu.service.DeleteMovie(ctx, id)
}
