package auth

import (
	auth "github.com/drakenchef/Tinder/internal/pkg/auth/repo"
)

type AuthUsecase struct {
	authRepo *auth.AuthRepo
}

func NewAuthUsecase(authRepo *auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{authRepo: authRepo}
}

func (au *AuthUsecase) Login() {
}

func (au *AuthUsecase) Logout() {
}
