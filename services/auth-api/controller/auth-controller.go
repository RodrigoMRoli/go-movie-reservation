package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/model"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/usecase"
)

type AuthController struct {
	usecase usecase.AuthUseCase
}

func NewAuthController(usecase usecase.AuthUseCase) AuthController {
	return AuthController{
		usecase: usecase,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var input model.LoginInputs
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := ac.usecase.Login(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
