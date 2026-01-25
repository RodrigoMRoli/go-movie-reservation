package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/controller"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/db"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/service"
	"github.com/rodrigomroli/go-movie-reservation/services/movie-api/usecase"
)

type Config struct {
	Port string `envconfig:"PORT"`
}

func main() {

	server := gin.Default()

	var cfg Config
	envErr := envconfig.Process("", &cfg)
	if envErr != nil {
		log.Fatal(envErr.Error())
	}

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
			"message": "Movies Service is up and running!",
		})
	})

	server.GET("/movies", MovieController.GetMovies)
	server.GET("/movies/:movieId", MovieController.GetMovie)
	server.POST("/movies", MovieController.CreateMovie)
	server.PATCH("/movies/:movieId", MovieController.UpdateMovie)
	server.DELETE("/movies/:movieId", MovieController.DeleteMovie)

	// Initialize Server
	server.Run(":" + cfg.Port)
}

func Indent(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", v)
	}
	return string(b)
}
