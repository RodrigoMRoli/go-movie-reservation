package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-movie-reservation/db"
	"go-movie-reservation/model"
	"go-movie-reservation/movie_resevation"
	"time"

	"github.com/google/uuid"
)

type MovieRepository struct {
	store   *db.SQLStore
	ctx     *context.Context
	queries *movie_resevation.Queries
}

func NewMovieRepository(store *db.SQLStore, ctx *context.Context, queries *movie_resevation.Queries) MovieRepository {
	return MovieRepository{
		store:   store,
		ctx:     ctx,
		queries: queries,
	}
}

func (mr *MovieRepository) GetMovies() ([]model.MovieWithGenre, error) {

	rows, err := mr.queries.GetMovies(*mr.ctx)
	if err != nil {
		fmt.Println(err)
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
				ID:            m.ID.String(),
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

func (mr *MovieRepository) GetMovie(id uuid.UUID) (model.MovieWithGenre, error) {

	m, err := mr.queries.GetMovie(*mr.ctx, id)
	if err != nil {
		fmt.Println(err)
		return model.MovieWithGenre{}, nil
	}

	var genres []string
	for _, g := range m.Genres {
		genres = append(genres, g)
	}
	movie := model.MovieWithGenre{
		Movie: model.Movie{
			ID:            m.ID.String(),
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
	params model.CreateMovieInput,
) (model.MovieWithGenre, error) {
	// deal with image management later
	// deal with transaction later if seem fit

	m, mErr := mr.queries.CreateMovie(*mr.ctx, movie_resevation.CreateMovieParams{
		Title: sql.NullString{
			String: params.Title,
			Valid:  true,
		},
		Description: sql.NullString{
			String: params.Description,
			Valid:  true,
		},
		PosterImage: sql.NullString{
			String: params.PosterImage,
			Valid:  true,
		},
		PosterExt: sql.NullString{
			String: params.PosterExt,
			Valid:  true,
		},
		ReleaseDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Language: sql.NullString{
			String: params.Language,
			Valid:  true,
		},
		CountryOrigin: sql.NullString{
			String: params.CountryOrigin,
			Valid:  true,
		},
	})

	if mErr != nil {
		fmt.Println(mErr)
		return model.MovieWithGenre{}, mErr
	}

	for _, genreId := range params.GenreIDs {
		genreMovieRow := movie_resevation.AddGenreToMovieParams{
			MovieID: uuid.NullUUID{
				UUID:  m.ID,
				Valid: true,
			},
			GenreID: uuid.NullUUID{
				UUID:  genreId,
				Valid: true,
			},
		}
		gErr := mr.queries.AddGenreToMovie(*mr.ctx, genreMovieRow)
		if gErr != nil {
			fmt.Println(gErr)
			return model.MovieWithGenre{}, gErr
		}
	}

	newMovie, err := mr.GetMovie(m.ID)
	if err != nil {
		fmt.Println(err)
		return model.MovieWithGenre{}, err
	}

	return newMovie, nil
}

func (mr *MovieRepository) UpdateMovie(
	id uuid.UUID,
	params model.CreateMovieInput,
) (model.MovieWithGenre, error) {

	args := movie_resevation.UpdateMovieParams{
		ID:            id,
		Title:         sql.NullString{String: params.Title, Valid: params.Title != ""},
		Description:   sql.NullString{String: params.Description, Valid: params.Description != ""},
		PosterImage:   sql.NullString{String: params.PosterImage, Valid: params.PosterImage != ""},
		PosterExt:     sql.NullString{String: params.PosterExt, Valid: params.PosterExt != ""},
		Minutes:       sql.NullInt32{Int32: int32(params.Minutes), Valid: params.Minutes != 0},
		ReleaseDate:   sql.NullTime{Time: params.ReleaseTime, Valid: !params.ReleaseTime.IsZero()},
		Language:      sql.NullString{String: params.Language, Valid: params.Language != ""},
		CountryOrigin: sql.NullString{String: params.CountryOrigin, Valid: params.CountryOrigin != ""},
	}

	err := mr.queries.UpdateMovie(*mr.ctx, args)
	if err != nil {
		fmt.Println(err)
		return model.MovieWithGenre{}, err
	}

	updatedMovie, err := mr.GetMovie(id)
	if err != nil {
		fmt.Println(err)
		return model.MovieWithGenre{}, err
	}

	return updatedMovie, nil
}

func (mr *MovieRepository) DeleteMovie(id uuid.UUID) error {
	err := mr.queries.DeleteMovie(*mr.ctx, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
