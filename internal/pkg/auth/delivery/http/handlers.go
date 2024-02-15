package auth

import (
	auth "github.com/drakenchef/Tinder/internal/pkg/auth/usecase"
	"net/http"
)

type AuthHandler struct {
	authUsecase *auth.AuthUsecase
}

func NewAuthHandler(authUsecase *auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (ah *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
}

func (ah *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
}

func (ah *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
}
