package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/controller"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/db"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/service"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/usecase"
)

func main() {

	server := gin.Default()

	database, dbErr := db.ConnectDB()
	if dbErr != nil {
		fmt.Println(dbErr)
	}

	defer database.Close()

	store := db.NewStore(database)

	// Services
	MoviceService := service.NewMovieService(store)

	// Usecases
	MovieUseCase := usecase.NewMovieUseCase(MoviceService)

	// Controllers
	MovieController := controller.NewMovieController(MovieUseCase)

	// Routes
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Everything is fine",
		})
	})

	server.GET("/movies", MovieController.GetMovies)
	server.GET("/movies/:movieId", MovieController.GetMovie)
	server.POST("/movies", MovieController.CreateMovie)
	server.PATCH("/movies/:movieId", MovieController.UpdateMovie)
	server.DELETE("/movies/:movieId", MovieController.DeleteMovie)

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
