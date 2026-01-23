package model

import "time"

type Movie struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	PosterImage   string    `json:"posterImage"`
	PosterExt     string    `json:"posterExt"`
	Minutes       int       `json:"minutes"`
	ReleaseDate   time.Time `json:"releaseDate"`
	Language      string    `json:"language"`
	CountryOrigin string    `json:"countryOrigin"`
}

type Genre struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type MovieWithGenre struct {
	Movie
	Genres []string `json:"genres"`
}
