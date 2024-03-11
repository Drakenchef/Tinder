package auth

import (
	"encoding/json"
	"fmt"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/auth"
	"github.com/drakenchef/Tinder/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

type AuthHandler struct {
	authUsecase auth.AuthUsecase
	logger      *zap.SugaredLogger
}

func NewAuthHandler(authUsecase auth.AuthUsecase, logger *zap.SugaredLogger) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
		logger:      logger,
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	var user models.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.authUsecase.CreateUser(r.Context(), user); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	var user models.SignInInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.authUsecase.GenerateToken(r.Context(), user)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
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
	utils.NameFuncLog()
	userId, err := utils.CheckAuth(r)
	fmt.Println(userId)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusForbidden)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
