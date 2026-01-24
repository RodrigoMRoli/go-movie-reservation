package main

import (
	"encoding/json"
	"fmt"
	"go-movie-reservation/controller"
	"go-movie-reservation/db"
	"go-movie-reservation/service"
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
