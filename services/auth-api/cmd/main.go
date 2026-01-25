package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/controller"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/db"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/service"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/usecase"
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

	// Store
	store := db.NewStore(database)

	// Services
	AuthService := service.NewAuthStore(store)

	// Usecases
	AuthUseCase := usecase.NewAuthUseCase(AuthService)

	// Controllers
	AuthController := controller.NewAuthController(AuthUseCase)

	// Routes
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Auth Service is up and running!",
		})
	})

	server.POST("/login", AuthController.Login)

	// Initialize Server
	server.Run(":" + cfg.Port)
}
