package controller

import (
	"go-movie-reservation/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type personController struct {
	personUseCase usecase.PersonUseCase
}

func NewPersonController(personUseCase usecase.PersonUseCase) personController {
	return personController{
		personUseCase: personUseCase,
	}
}

func (pc *personController) GetPeople(ctx *gin.Context) {
	people, err := pc.personUseCase.GetPeople()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, people)
}
