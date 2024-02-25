package auth

import (
	"encoding/json"
	"fmt"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/auth"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/pkg/errors"
	"net/http"
)

type AuthHandler struct {
	authUsecase auth.AuthUsecase
}

func NewAuthHandler(authUsecase auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.SignInInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to decode request body").Error(), http.StatusBadRequest)
		return
	}

	err = h.authUsecase.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to sign up user").Error(), http.StatusInternalServerError)
		return
	}
	//commit
	w.WriteHeader(http.StatusCreated)
}
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.SignInInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to decode request body").Error(), http.StatusBadRequest)
		return
	}

	token, err := h.authUsecase.GenerateToken(r.Context(), user)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to sign in user").Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.CheckAuth(r)
	fmt.Println(userId)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("User is not authorized"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	} else {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("User is authorized"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}
