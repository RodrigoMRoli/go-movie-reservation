package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-movie-reservation/controller"
	"go-movie-reservation/db"
	"go-movie-reservation/movie_resevation"
	"go-movie-reservation/repository"
	"go-movie-reservation/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	database, dbErr := db.ConnectDB()
	if dbErr != nil {
		fmt.Println(dbErr)
	}

	defer database.Close()

	store := db.NewStore(database)
	ctx := context.Background()
	queries := movie_resevation.New(database)

	// Repositories
	MovieRepository := repository.NewMovieRepository(store, &ctx, queries)

	// Usecases
	MovieUseCase := usecase.NewMovieUseCase(MovieRepository)

	// Controllers
	MovieController := controller.NewMovieController(MovieUseCase)

	// Routes
	server.GET("/health", func(ginCtx *gin.Context) {
		ginCtx.JSON(200, gin.H{
			"message": "Everything is fine",
		})
	})

	server.GET("/movies", MovieController.GetMovies)
	server.GET("/movies/:movieId", MovieController.GetMovie)
	server.POST("/movies", MovieController.CreateMovie)
	server.PATCH("/movies/:movieId", MovieController.UpdateMovie)

	// Initialize Server
	server.Run(":8080")
}

func Indent(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", v)
	}
	return string(b)
}
