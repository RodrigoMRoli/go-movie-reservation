package usecase

import (
	"go-movie-reservation/model"
	"go-movie-reservation/repository"
)

type MovieUseCase struct {
	repository repository.MovieRepository
}

func NewMovieUseCase(repository repository.MovieRepository) MovieUseCase {
	return MovieUseCase{
		repository: repository,
	}
}

func (mu *MovieUseCase) GetMovies() ([]model.Movie, error) {
	return mu.repository.GetMovies()
}
