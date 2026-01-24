package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-movie-reservation/movie_resevation"
	"go-movie-reservation/repository"
)

type MockQuerier struct {
	mock.Mock
}

func (m *MockQuerier) GetMovie(ctx context.Context, id uuid.UUID) (movie_resevation.GetMovieRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(movie_resevation.GetMovieRow), args.Error(1)
}
func (m *MockQuerier) AddGenreToMovie(ctx context.Context, arg movie_resevation.AddGenreToMovieParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}
func (m *MockQuerier) CreateGenre(ctx context.Context, title sql.NullString) (movie_resevation.MvGenre, error) {
	args := m.Called(ctx, title)
	return args.Get(0).(movie_resevation.MvGenre), args.Error(1)
}
func (m *MockQuerier) CreateMovie(ctx context.Context, arg movie_resevation.CreateMovieParams) (movie_resevation.MvMovie, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(movie_resevation.MvMovie), args.Error(1)
}
func (m *MockQuerier) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockQuerier) DeleteGenre(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockQuerier) GetMovies(ctx context.Context) ([]movie_resevation.GetMoviesRow, error) {
	args := m.Called(ctx)
	return args.Get(0).([]movie_resevation.GetMoviesRow), args.Error(1)
}
func (m *MockQuerier) RemoveGenreFromMovie(ctx context.Context, arg movie_resevation.RemoveGenreFromMovieParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}
func (m *MockQuerier) UpdateMovie(ctx context.Context, arg movie_resevation.UpdateMovieParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func TestGetMovie_Success(t *testing.T) {
	// Arrange
	mockQuerier := new(MockQuerier)

	repo := repository.NewMovieRepository(mockQuerier)

	movieID := uuid.New()

	// Mocked Data
	expectedDBMovie := movie_resevation.GetMovieRow{
		ID:          movieID,
		Title:       sql.NullString{String: "Inception", Valid: true},
		Description: sql.NullString{String: "Dream within a dream", Valid: true},
		Minutes:     sql.NullInt32{Int32: 148, Valid: true},
	}

	mockQuerier.On("GetMovie", mock.Anything, movieID).Return(expectedDBMovie, nil)

	// Act
	ctx := context.Background()
	result, err := repo.GetMovie(ctx, movieID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Inception", result.Movie.Title)
	assert.Equal(t, movieID.String(), result.Movie.ID)

	mockQuerier.AssertExpectations(t)
}
