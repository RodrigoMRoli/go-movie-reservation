package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/model"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/usecase"
)

type movieController struct {
	movieUseCase usecase.MovieUseCase
}

func NewMovieController(movieUseCase usecase.MovieUseCase) movieController {
	return movieController{
		movieUseCase: movieUseCase,
	}
}

func (mc *movieController) GetMovies(c *gin.Context) {
	ctx := c.Request.Context()
	movies, err := mc.movieUseCase.GetMovies(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (mc *movieController) GetMovie(c *gin.Context) {
	ctx := c.Request.Context()
	paramsId := c.Param("movieId")
	if paramsId == "" {
		response := model.Response{
			Message: "Id cannot be null",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := uuid.Parse(paramsId)
	if err != nil {
		response := model.Response{
			Message: "Id is invalid",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	movie, err := mc.movieUseCase.GetMovie(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (mc *movieController) CreateMovie(c *gin.Context) {
	ctx := c.Request.Context()
	var input model.CreateMovieInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newMovie, err := mc.movieUseCase.CreateMovie(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newMovie)
}

func (mc *movieController) UpdateMovie(c *gin.Context) {
	ctx := c.Request.Context()
	paramsId := c.Param("movieId")
	if paramsId == "" {
		response := model.Response{
			Message: "Id cannot be null",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id, uuidErr := uuid.Parse(paramsId)
	if uuidErr != nil {
		response := model.Response{
			Message: "Id is invalid",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input model.UpdateMovieInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newMovie, err := mc.movieUseCase.UpdateMovie(ctx, id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newMovie)
}

func (mc *movieController) DeleteMovie(c *gin.Context) {
	ctx := c.Request.Context()
	paramsId := c.Param("movieId")
	if paramsId == "" {
		response := model.Response{
			Message: "Id cannot be null",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := uuid.Parse(paramsId)
	if err != nil {
		response := model.Response{
			Message: "Id is invalid",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteErr := mc.movieUseCase.DeleteMovie(ctx, id)
	if deleteErr != nil {
		c.JSON(http.StatusInternalServerError, deleteErr)
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "Movie deleted successfully",
	})
}
