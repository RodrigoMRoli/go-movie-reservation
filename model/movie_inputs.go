package model

import (
	"time"
)

type MovieInput struct {
	Title         string
	Description   string
	PosterImage   string
	PosterExt     string
	Minutes       int
	ReleaseTime   time.Time
	Language      string
	CountryOrigin string
}

type CreateMovieInput struct {
	MovieInput
	Genres []string
}

type UpdateMovieInput struct {
	MovieInput
	AddGenres    []string
	RemoveGenres []string
}
