package controller

import (
	"go-movie-reservation/model"
	"go-movie-reservation/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func (mc *movieController) GetMovie(ctx *gin.Context) {
	ctxId := ctx.Param("movieId")
	if ctxId == "" {
		response := model.Response{
			Message: "Id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := uuid.Parse(ctxId)
	if err != nil {
		response := model.Response{
			Message: "Id is invalid",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	movie, err := mc.movieUseCase.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (mc *movieController) CreateMovie(ctx *gin.Context) {

	var input model.CreateMovieInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newMovie, err := mc.movieUseCase.CreateMovie(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, newMovie)
}

func (mc *movieController) UpdateMovie(ctx *gin.Context) {
	ctxId := ctx.Param("movieId")
	if ctxId == "" {
		response := model.Response{
			Message: "Id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, uuidErr := uuid.Parse(ctxId)
	if uuidErr != nil {
		response := model.Response{
			Message: "Id is invalid",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var input model.CreateMovieInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newMovie, err := mc.movieUseCase.UpdateMovie(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, newMovie)
}

func (mc *movieController) DeleteMovie(ctx *gin.Context) {
	ctxId := ctx.Param("movieId")
	if ctxId == "" {
		response := model.Response{
			Message: "Id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := uuid.Parse(ctxId)
	if err != nil {
		response := model.Response{
			Message: "Id is invalid",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deleteErr := mc.movieUseCase.DeleteMovie(id)
	if deleteErr != nil {
		ctx.JSON(http.StatusInternalServerError, deleteErr)
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Movie deleted successfully",
	})
}
