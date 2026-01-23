package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateMovieInput struct {
	Title         string
	Description   string
	PosterImage   string
	PosterExt     string
	Minutes       int
	ReleaseTime   time.Time
	Language      string
	CountryOrigin string
	GenreIDs      []uuid.UUID
}
