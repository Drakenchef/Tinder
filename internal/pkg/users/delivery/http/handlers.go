package users

import (
	"context"
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/middleware/csrfmw"
	"github.com/drakenchef/Tinder/internal/pkg/users"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type UsersHandler struct {
	usersUsecase users.UsersUsecase
	logger       *zap.SugaredLogger
}

const maxRequestBodySize = 1024 * 1024 * 5

func NewUsersHandler(usersUsecase users.UsersUsecase, logger *zap.SugaredLogger) *UsersHandler {
	return &UsersHandler{
		usersUsecase: usersUsecase,
		logger:       logger,
	}
}

func (h *UsersHandler) UsersList(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	ctx := context.WithValue(r.Context(), "id", r.Header.Get("abcd"))
	//h.logger.Info(ctx.Value("id"))
	//h.logger.Info("UsersListHandler")
	users, err := h.usersUsecase.UsersList(ctx)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to sign up user").Error(), http.StatusInternalServerError)
		return
	}
	//log.Print(r.Header.Get("abcd"))
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func (h *UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	var req struct {
		UID string `json:"uid"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(req.UID)
	if err != nil {
		http.Error(w, "invalid user ID format", http.StatusBadRequest)
		return
	}

	user, err := h.usersUsecase.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	uidFromContext := r.Header.Get("uid")
	h.logger.Info("get uid from context: ", uidFromContext)
	uid, _ := uuid.Parse(uidFromContext)

	tokenExpTime := int64(3600)

	csrf, _ := csrfmw.NewHMACKHashToken("zxczxczczxc", h.logger)
	csrfToken, err := csrf.Create(uid, tokenExpTime)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("X-CSRF-Token", csrfToken)
	userJSON, err := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)

}

func (h *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.usersUsecase.UpdateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	userJSON, err := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)

}

func (h *UsersHandler) UpdateUserImage(w http.ResponseWriter, r *http.Request) {
	uidFromContext := r.Header.Get("uid")
	h.logger.Info("get uid from context: ", uidFromContext)
	uid, _ := uuid.Parse(uidFromContext)

	limitedReader := http.MaxBytesReader(w, r.Body, maxRequestBodySize)
	defer r.Body.Close()

	bodyContent, err := io.ReadAll(limitedReader)
	fileFormat := http.DetectContentType(bodyContent)
	if err != nil && !errors.Is(err, io.EOF) {
		if errors.As(err, new(*http.MaxBytesError)) {
			w.WriteHeader(http.StatusRequestEntityTooLarge)

		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			h.logger.Info("error: %s", err)
		}

		return
	}

	user, err := h.usersUsecase.UpdateUserImage(r.Context(), uid, bodyContent, fileFormat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Info("error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	userJSON, _ := json.Marshal(user)
	w.Write(userJSON)
}
func (h *UsersHandler) DeleteUserImage(w http.ResponseWriter, r *http.Request) {
	uidFromContext := r.Header.Get("uid")
	h.logger.Info("get uid from context: ", uidFromContext)
	uid, _ := uuid.Parse(uidFromContext)
	var req struct {
		URL string `json:"url" binding:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.usersUsecase.DeleteUserImage(r.Context(), req.URL, uid); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *UsersHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	uidFromContext := r.Header.Get("uid")
	h.logger.Info("get uid from context: ", uidFromContext)
	uid, _ := uuid.Parse(uidFromContext)

	var request models.ChangePassword
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.usersUsecase.UpdateUserPassword(r.Context(), request, uid)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	uidFromContext := r.Header.Get("uid")
	h.logger.Info("get uid from context: ", uidFromContext)
	uid, _ := uuid.Parse(uidFromContext)

	var password models.ChangePassword
	if err := json.NewDecoder(r.Body).Decode(&password); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.usersUsecase.DeleteUser(r.Context(), password, uid); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
