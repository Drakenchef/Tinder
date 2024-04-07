package auth

import (
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/auth/delivery/grpc/gen"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

//	type AuthHandler struct {
//		authUsecase auth.AuthUsecase
//		logger      *zap.SugaredLogger
//	}
type AuthHandler struct {
	client gen.AuthClient
	logger *zap.SugaredLogger
}

//func NewAuthHandler(authUsecase auth.AuthUsecase, logger *zap.SugaredLogger) *AuthHandler {
//	return &AuthHandler{
//		authUsecase: authUsecase,
//		logger:      logger,
//	}
//}

func NewAuthHandler(cl gen.AuthClient, logger *zap.SugaredLogger) *AuthHandler {
	return &AuthHandler{
		client: cl,
		logger: logger,
	}
}
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	userInfo := &models.SignInInput{}
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.client.SignUp(r.Context(), &gen.SignUpRequest{
		Login:    userInfo.Login,
		Password: userInfo.Password,
	})
	if err != nil {
		h.logger.Info("failed to signup", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	userInfo := &models.SignInInput{}
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := h.client.SignIn(r.Context(), &gen.SignInRequest{
		Login:    userInfo.Login,
		Password: userInfo.Password,
	})
	if err != nil {
		h.logger.Info("failed to signin", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Authorization", "Bearer "+response.Token)
	cookie := http.Cookie{
		Name:     "token",
		Value:    response.Token,
		Path:     "/",
		HttpOnly: true,
		//Secure:   true,
		//SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

//
//func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
//	utils.NameFuncLog()
//	var user models.SignInInput
//	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//		h.logger.Info(err)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	if err := h.authUsecase.CreateUser(r.Context(), user); err != nil {
//		h.logger.Info(err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusCreated)
//}
//
//func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
//	utils.NameFuncLog()
//	var user models.SignInInput
//	err := json.NewDecoder(r.Body).Decode(&user)
//	if err != nil {
//		h.logger.Info(err)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	token, err := h.authUsecase.GenerateToken(r.Context(), user)
//	if err != nil {
//		h.logger.Info(err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Authorization", "Bearer "+token)
//	cookie := http.Cookie{
//		Name:     "token",
//		Value:    token,
//		Path:     "/",
//		HttpOnly: true,
//		//Secure:   true,
//		//SameSite: http.SameSiteStrictMode,
//	}
//	http.SetCookie(w, &cookie)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	uidFromContext := r.Header.Get("uid")
	h.logger.Info("get uid from context: ", uidFromContext)
	uid, _ := uuid.Parse(uidFromContext)
	if uid == uuid.Nil {
		w.WriteHeader(http.StatusForbidden)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
