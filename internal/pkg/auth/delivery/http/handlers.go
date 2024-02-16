package auth

import (
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/models"
	auth "github.com/drakenchef/Tinder/internal/pkg/auth/usecase"
	"github.com/pkg/errors"
	"net/http"
)

type AuthHandler struct {
	authUsecase *auth.AuthUsecase
}

func NewAuthHandler(authUsecase *auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to decode request body").Error(), http.StatusBadRequest)
		return
	}

	err = h.authUsecase.CreateUser(user)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to sign up user").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to decode request body").Error(), http.StatusBadRequest)
		return
	}

	token, err := h.authUsecase.GenerateToken(user.Login, user.Password)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to sign in user").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func (ah *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
}
