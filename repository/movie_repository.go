package repository

import (
	"context"
	"fmt"
	"go-movie-reservation/model"
	"go-movie-reservation/movie_resevation"
)

type MovieRepository struct {
	ctx     *context.Context
	queries *movie_resevation.Queries
}

func NewMovieRepository(ctx *context.Context, queries *movie_resevation.Queries) MovieRepository {
	return MovieRepository{
		ctx:     ctx,
		queries: queries,
	}
}

func (mr *MovieRepository) GetMovies() ([]model.Movie, error) {

	rows, err := mr.queries.GetMovies(*mr.ctx)
	if err != nil {
		fmt.Println(err)
		return []model.Movie{}, nil
	}

	var movies []model.Movie
	for _, m := range rows {
		movies = append(movies, model.Movie{
			ID:            m.ID.String(),
			Title:         m.Title.String,
			Description:   m.Description.String,
			PosterImage:   m.PosterImage.String,
			Minutes:       int(m.Minutes.Int32),
			ReleaseDate:   m.ReleaseDate.Time.Format("DateOnly"),
			Language:      m.Language.String,
			CountryOrigin: m.CountryOrigin.String,
		})
	}

	return movies, nil
}
