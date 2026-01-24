package model

import (
	"time"
)

type MovieInput struct {
	Title         *string    `json:"title"`
	Description   *string    `json:"description"`
	PosterImage   *string    `json:"poster_image"`
	PosterExt     *string    `json:"poster_ext"`
	Minutes       *int       `json:"minutes"`
	ReleaseDate   *time.Time `json:"release_time"`
	Language      *string    `json:"language"`
	CountryOrigin *string    `json:"country_origin"`
}

type CreateMovieInput struct {
	MovieInput
	Genres []string `json:"genres"`
}

type UpdateMovieInput struct {
	MovieInput
	AddGenres    []string `json:"addGenres"`
	RemoveGenres []string `json:"removeGenres"`
}
