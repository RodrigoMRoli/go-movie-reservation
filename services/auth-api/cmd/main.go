package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/db"
)

func main() {
	server := gin.Default()

	database, dbErr := db.ConnectDB()
	if dbErr != nil {
		fmt.Println(dbErr)
	}

	defer database.Close()

	// Store
	// store := db.NewStore(database)

	// Services

	// Usecases

	// Controllers

	// Routes
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Auth Service is up and running!",
		})
	})

	// Initialize Server
	server.Run(":8080")
}
