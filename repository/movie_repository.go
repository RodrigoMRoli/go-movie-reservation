package repository

import (
	"context"
	"database/sql"
	"go-movie-reservation/helpers"
	"go-movie-reservation/model"
	"go-movie-reservation/movie_resevation"

	"github.com/google/uuid"
)

type MovieRepository struct {
	db      *sql.DB
	queries movie_resevation.Queries
}

func NewMovieRepository(db *sql.DB, queries movie_resevation.Queries) MovieRepository {
	return MovieRepository{
		queries: queries,
	}
}

func (mr *MovieRepository) GetMovies(ctx context.Context) ([]model.MovieWithGenre, error) {

	rows, err := mr.queries.GetMovies(ctx)
	if err != nil {
		return []model.MovieWithGenre{}, nil
	}

	var movies []model.MovieWithGenre
	for _, m := range rows {

		var genres []string
		for _, g := range m.Genres {
			genres = append(genres, g)
		}

		movie := model.MovieWithGenre{
			Movie: model.Movie{
				ID:            m.ID,
				Title:         m.Title.String,
				Description:   m.Description.String,
				PosterImage:   m.PosterImage.String,
				Minutes:       int(m.Minutes.Int32),
				ReleaseDate:   m.ReleaseDate.Time,
				Language:      m.Language.String,
				CountryOrigin: m.CountryOrigin.String,
			},
			Genres: genres,
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (mr *MovieRepository) GetMovie(ctx context.Context, id uuid.UUID) (model.MovieWithGenre, error) {

	m, err := mr.queries.GetMovie(ctx, id)
	if err != nil {
		return model.MovieWithGenre{}, nil
	}

	var genres []string
	for _, g := range m.Genres {
		genres = append(genres, g)
	}
	movie := model.MovieWithGenre{
		Movie: model.Movie{
			ID:            m.ID,
			Title:         m.Title.String,
			Description:   m.Description.String,
			PosterImage:   m.PosterImage.String,
			PosterExt:     m.PosterExt.String,
			Minutes:       int(m.Minutes.Int32),
			ReleaseDate:   m.ReleaseDate.Time,
			Language:      m.Language.String,
			CountryOrigin: m.CountryOrigin.String,
		},
		Genres: genres,
	}

	return movie, nil
}

func (mr *MovieRepository) CreateMovie(
	ctx context.Context,
	params model.CreateMovieInput,
) (model.MovieWithGenre, error) {
	// deal with image management later
	// deal with transaction later if seem fit

	tx, err := mr.db.BeginTx(ctx, nil)
	if err != nil {
		return model.MovieWithGenre{}, err
	}
	defer tx.Rollback()

	qtx := mr.queries.WithTx(tx)

	args := movie_resevation.CreateMovieParams{
		Title:         helpers.StringPointerToNullString(params.Title),
		Description:   helpers.StringPointerToNullString(params.Description),
		PosterImage:   helpers.StringPointerToNullString(params.PosterImage),
		PosterExt:     helpers.StringPointerToNullString(params.PosterExt),
		Minutes:       helpers.IntPointerToNullInt32(params.Minutes),
		ReleaseDate:   helpers.TimePointerToNullTime(params.ReleaseDate),
		Language:      helpers.StringPointerToNullString(params.Language),
		CountryOrigin: helpers.StringPointerToNullString(params.CountryOrigin),
	}

	m, mErr := qtx.CreateMovie(ctx, args)

	if mErr != nil {
		return model.MovieWithGenre{}, err
	}

	for _, genre := range params.Genres {
		genreMovieRow := movie_resevation.AddGenreToMovieParams{
			MovieID: uuid.NullUUID{
				UUID:  m.ID,
				Valid: true,
			},
			Title: sql.NullString{
				String: genre,
				Valid:  true,
			},
		}
		gErr := qtx.AddGenreToMovie(ctx, genreMovieRow)
		if gErr != nil {
			return model.MovieWithGenre{}, gErr
		}
	}

	newMovie, err := mr.GetMovie(ctx, m.ID)
	if err != nil {
		return model.MovieWithGenre{}, err
	}

	return newMovie, nil
}

func (mr *MovieRepository) UpdateMovie(
	ctx context.Context,
	id uuid.UUID,
	params model.UpdateMovieInput,
) (model.MovieWithGenre, error) {

	args := movie_resevation.UpdateMovieParams{
		ID:            id,
		Title:         helpers.StringPointerToNullString(params.Title),
		Description:   helpers.StringPointerToNullString(params.Description),
		PosterImage:   helpers.StringPointerToNullString(params.PosterImage),
		PosterExt:     helpers.StringPointerToNullString(params.PosterExt),
		Minutes:       helpers.IntPointerToNullInt32(params.Minutes),
		ReleaseDate:   helpers.TimePointerToNullTime(params.ReleaseDate),
		Language:      helpers.StringPointerToNullString(params.Language),
		CountryOrigin: helpers.StringPointerToNullString(params.CountryOrigin),
	}

	// Update Movie
	err := mr.queries.UpdateMovie(ctx, args)
	if err != nil {
		return model.MovieWithGenre{}, err
	}

	// Add Genres
	if len(params.AddGenres) > 0 {
		for _, genre := range params.AddGenres {
			addGenreRow := movie_resevation.AddGenreToMovieParams{
				MovieID: uuid.NullUUID{
					UUID:  id,
					Valid: true,
				},
				Title: sql.NullString{
					String: genre,
					Valid:  true,
				},
			}
			gErr := mr.queries.AddGenreToMovie(ctx, addGenreRow)
			if gErr != nil {
				return model.MovieWithGenre{}, gErr
			}
		}
	}

	// Remove Genres
	if len(params.RemoveGenres) > 0 {
		for _, genre := range params.RemoveGenres {
			removeGenreRow := movie_resevation.RemoveGenreFromMovieParams{
				MovieID: uuid.NullUUID{
					UUID:  id,
					Valid: true,
				},
				Title: sql.NullString{
					String: genre,
					Valid:  true,
				},
			}
			gErr := mr.queries.RemoveGenreFromMovie(ctx, removeGenreRow)
			if gErr != nil {
				return model.MovieWithGenre{}, gErr
			}
		}
	}

	updatedMovie, err := mr.GetMovie(ctx, id)
	if err != nil {
		return model.MovieWithGenre{}, err
	}

	return updatedMovie, nil
}

func (mr *MovieRepository) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	err := mr.queries.DeleteMovie(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
