package service

import (
	"context"
	"database/sql"
	"go-movie-reservation/db"
	"go-movie-reservation/helpers"
	"go-movie-reservation/model"
	"go-movie-reservation/movie_resevation"

	"github.com/google/uuid"
)

type MovieService struct {
	store db.Store
}

func NewMovieService(store db.Store) MovieService {
	return MovieService{store: store}
}

func (ms *MovieService) GetMovies(ctx context.Context) ([]model.MovieWithGenre, error) {

	rows, err := ms.store.GetMovies(ctx)
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

func (ms *MovieService) GetMovie(ctx context.Context, id uuid.UUID) (model.MovieWithGenre, error) {

	m, err := ms.store.GetMovie(ctx, id)
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

func (ms *MovieService) CreateMovie(
	ctx context.Context,
	params model.CreateMovieInput,
) (model.MovieWithGenre, error) {

	var resultID uuid.UUID

	err := ms.store.ExecTx(ctx, func(q movie_resevation.Querier) error {

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

		m, err := q.CreateMovie(ctx, args)
		if err != nil {
			return err
		}

		resultID = m.ID

		for _, genre := range params.Genres {
			genreMovieRow := movie_resevation.AddGenreToMovieParams{
				MovieID: uuid.NullUUID{UUID: m.ID, Valid: true},
				Title:   sql.NullString{String: genre, Valid: true},
			}
			if err := q.AddGenreToMovie(ctx, genreMovieRow); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return model.MovieWithGenre{}, err
	}

	newMovie := model.MovieWithGenre{
		Movie: model.Movie{
			ID:            resultID,
			Title:         *params.Title,
			Description:   *params.Description,
			PosterImage:   *params.PosterImage,
			PosterExt:     *params.PosterExt,
			Minutes:       *params.Minutes,
			ReleaseDate:   *params.ReleaseDate,
			Language:      *params.Language,
			CountryOrigin: *params.CountryOrigin,
		},
		Genres: params.Genres,
	}

	return newMovie, nil
}

func (ms *MovieService) UpdateMovie(
	ctx context.Context,
	id uuid.UUID,
	params model.UpdateMovieInput,
) (model.MovieWithGenre, error) {

	var movie model.MovieWithGenre

	txErr := ms.store.ExecTx(ctx, func(tx movie_resevation.Querier) error {

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
		err := tx.UpdateMovie(ctx, args)
		if err != nil {
			return err
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
				err := tx.AddGenreToMovie(ctx, addGenreRow)
				if err != nil {
					return err
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
				err := tx.RemoveGenreFromMovie(ctx, removeGenreRow)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if txErr != nil {
		return model.MovieWithGenre{}, txErr
	}

	movie, err := ms.GetMovie(ctx, id)
	if err != nil {
		return model.MovieWithGenre{}, txErr
	}

	return movie, nil
}

func (ms *MovieService) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	err := ms.store.DeleteMovie(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
