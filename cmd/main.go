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

	db, dbErr := db.ConnectDB()
	if dbErr != nil {
		fmt.Println(dbErr)
	}

	defer db.Close()

	ctx := context.Background()
	queries := movie_resevation.New(db)

	// Repositories
	PersonRepository := repository.NewPersonRepository(&ctx, queries)

	// Usecases
	PersonUseCase := usecase.NewPersonUseCase(PersonRepository)

	// Controllers
	PersonController := controller.NewPersonController(PersonUseCase)

	// Routes
	server.GET("/health", func(ginCtx *gin.Context) {
		ginCtx.JSON(200, gin.H{
			"message": "Everything is fine",
		})
	})

	server.GET("/people", PersonController.GetPeople)

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
