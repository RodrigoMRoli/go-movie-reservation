package usecase

import (
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

func (mu *MovieUseCase) GetMovies() ([]model.MovieWithGenre, error) {
	return mu.repository.GetMovies()
}

func (mu *MovieUseCase) GetMovie(id uuid.UUID) (model.MovieWithGenre, error) {
	return mu.repository.GetMovie(id)
}

func (mu *MovieUseCase) CreateMovie(params model.CreateMovieInput) (model.MovieWithGenre, error) {
	return mu.repository.CreateMovie(params)
}

func (mu *MovieUseCase) UpdateMovie(id uuid.UUID, params model.CreateMovieInput) (model.MovieWithGenre, error) {
	return mu.repository.UpdateMovie(id, params)
}

func (mu *MovieUseCase) DeleteMovie(id uuid.UUID) error {
	return mu.repository.DeleteMovie(id)
}
