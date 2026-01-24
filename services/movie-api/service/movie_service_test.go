package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/model"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/movie_resevation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
	movie_resevation.Querier
}

func (m *MockStore) ExecTx(ctx context.Context, fn func(movie_resevation.Querier) error) error {
	return fn(m)
}

func (m *MockStore) CreateMovie(ctx context.Context, arg movie_resevation.CreateMovieParams) (movie_resevation.MvMovie, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(movie_resevation.MvMovie), args.Error(1)
}

func (m *MockStore) UpdateMovie(ctx context.Context, arg movie_resevation.UpdateMovieParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockStore) AddGenreToMovie(ctx context.Context, arg movie_resevation.AddGenreToMovieParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockStore) RemoveGenreFromMovie(ctx context.Context, arg movie_resevation.RemoveGenreFromMovieParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockStore) GetMovie(ctx context.Context, id uuid.UUID) (movie_resevation.GetMovieRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(movie_resevation.GetMovieRow), args.Error(1)
}

func (m *MockStore) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockStore) GetMovies(ctx context.Context) ([]movie_resevation.GetMoviesRow, error) {
	args := m.Called(ctx)
	return args.Get(0).([]movie_resevation.GetMoviesRow), args.Error(1)
}

func (m *MockStore) CreateGenre(ctx context.Context, name sql.NullString) (movie_resevation.MvGenre, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(movie_resevation.MvGenre), args.Error(1)
}

func TestUpdateMovie_Success(t *testing.T) {
	// Arrange
	mockStore := &MockStore{}
	service := NewMovieService(mockStore)
	ctx := context.Background()
	id := uuid.New()

	newTitle := "Updated Title"
	newDesc := "Updated Desc"
	params := model.UpdateMovieInput{
		MovieInput: model.MovieInput{
			Title:       &newTitle,
			Description: &newDesc,
		},
		AddGenres:    []string{"Drama"},
		RemoveGenres: []string{"Comedy"},
	}

	mockStore.On("UpdateMovie", ctx, mock.MatchedBy(func(arg movie_resevation.UpdateMovieParams) bool {
		return arg.ID == id && arg.Title.String == "Updated Title" && arg.Title.Valid == true
	})).Return(nil)

	mockStore.On("AddGenreToMovie", ctx, mock.MatchedBy(func(arg movie_resevation.AddGenreToMovieParams) bool {
		return arg.MovieID.UUID == id && arg.Title.String == "Drama"
	})).Return(nil)

	mockStore.On("RemoveGenreFromMovie", ctx, mock.MatchedBy(func(arg movie_resevation.RemoveGenreFromMovieParams) bool {
		return arg.MovieID.UUID == id && arg.Title.String == "Comedy"
	})).Return(nil)

	mockResponse := movie_resevation.GetMovieRow{
		ID:          id,
		Title:       sql.NullString{String: "Updated Title", Valid: true},
		Description: sql.NullString{String: "Updated Desc", Valid: true},
		Genres:      []string{"Drama"},
		Minutes:     sql.NullInt32{Int32: 120, Valid: true},
		ReleaseDate: sql.NullTime{Time: time.Now(), Valid: true},
	}
	mockStore.On("GetMovie", ctx, id).Return(mockResponse, nil)

	// Act
	result, err := service.UpdateMovie(ctx, id, params)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", result.Movie.Title)
	assert.Equal(t, "Updated Desc", result.Movie.Description)
	assert.Contains(t, result.Genres, "Drama")

	mockStore.AssertExpectations(t)
}

func TestUpdateMovie_TransactionError(t *testing.T) {
	// Arrange
	mockStore := new(MockStore)
	service := NewMovieService(mockStore)
	ctx := context.Background()
	id := uuid.New()

	params := model.UpdateMovieInput{
		AddGenres: []string{"Horror"},
	}

	mockStore.On("UpdateMovie", ctx, mock.Anything).Return(nil) // Update Ok

	expectedErr := errors.New("database connection lost") // AddGenre fails
	mockStore.On("AddGenreToMovie", ctx, mock.Anything).Return(expectedErr)

	// Act
	result, err := service.UpdateMovie(ctx, id, params)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, uuid.Nil, result.Movie.ID)

	mockStore.AssertNotCalled(t, "GetMovie")
	mockStore.AssertExpectations(t)
}

func TestCreateMovie_Success(t *testing.T) {
	// Arrange
	mockStore := &MockStore{}
	service := NewMovieService(mockStore)
	ctx := context.Background()

	title := "Inception"
	params := model.CreateMovieInput{
		MovieInput: model.MovieInput{
			Title: &title,
		},
		Genres: []string{"Sci-Fi"},
	}

	generatedID := uuid.New()

	mockStore.On("CreateMovie", ctx, mock.MatchedBy(func(arg movie_resevation.CreateMovieParams) bool {
		return arg.Title.String == "Inception"
	})).Return(movie_resevation.MvMovie{ID: generatedID}, nil)

	mockStore.On("AddGenreToMovie", ctx, mock.Anything).Return(nil)

	mockStore.On("GetMovie", ctx, generatedID).Return(movie_resevation.GetMovieRow{
		ID:     generatedID,
		Title:  sql.NullString{String: "Inception", Valid: true},
		Genres: []string{"Sci-Fi"},
	}, nil)

	// Act
	result, err := service.CreateMovie(ctx, params)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, generatedID, result.Movie.ID)
	assert.Equal(t, "Inception", result.Movie.Title)

	mockStore.AssertExpectations(t)
}
