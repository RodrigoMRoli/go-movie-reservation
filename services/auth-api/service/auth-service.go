package service

import (
	"context"

	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/auth"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/db"
	"github.com/rodrigomroli/go-movie-reservation/services/auth-api/model"
)

type AuthService struct {
	store db.Store
}

func NewAuthStore(store db.Store) AuthService {
	return AuthService{store: store}
}

func (as AuthService) Login(ctx context.Context, input model.LoginInputs) (auth.User, error) {
	user, err := as.store.GetUser(ctx, input.Email)
	if err != nil {
		return auth.User{}, err
	}

	return user, nil
}
