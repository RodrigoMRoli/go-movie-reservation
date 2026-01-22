package controller

import (
	"go-movie-reservation/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type movieController struct {
	movieUseCase usecase.MovieUseCase
}

func NewMovieController(movieUseCase usecase.MovieUseCase) movieController {
	return movieController{
		movieUseCase: movieUseCase,
	}
}

func (mc *movieController) GetMovies(ctx *gin.Context) {
	movies, err := mc.movieUseCase.GetMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, movies)
}
