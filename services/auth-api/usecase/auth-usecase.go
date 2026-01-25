package usecase

import (
	"context"

	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/auth"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/model"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/service"
)

type AuthUseCase struct {
	service service.AuthService
}

func NewAuthUseCase(service service.AuthService) AuthUseCase {
	return AuthUseCase{
		service: service,
	}
}

func (au *AuthUseCase) Login(ctx context.Context, input model.LoginInputs) (auth.User, error) {
	return au.service.Login(ctx, input)
}
