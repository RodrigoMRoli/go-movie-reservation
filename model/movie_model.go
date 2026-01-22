package model

type Movie struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	PosterImage   string `json:"posterImage"`
	Minutes       int    `json:"minutes"`
	ReleaseDate   string `json:"releaseDate"`
	Language      string `json:"language"`
	CountryOrigin string `json:"countryOrigin"`
}
