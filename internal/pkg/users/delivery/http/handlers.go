package users

import (
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/users"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type UsersHandler struct {
	usersUsecase users.UsersUsecase
}

const maxRequestBodySize = 1024 * 1024 * 5

func NewUsersHandler(usersUsecase users.UsersUsecase) *UsersHandler {
	return &UsersHandler{usersUsecase: usersUsecase}
}

func (h *UsersHandler) UsersList(w http.ResponseWriter, r *http.Request) {
	users, err := h.usersUsecase.UsersList(r.Context())
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to sign up user").Error(), http.StatusInternalServerError)
		return
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func (h *UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
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
	userJSON, err := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)

}

func (h *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No cookie present", http.StatusUnauthorized)
		} else {
			http.Error(w, "Error retrieving cookie", http.StatusBadRequest)
		}
		return
	}
	token := cookie.Value
	uid, err := utils.ParseToken(token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	limitedReader := http.MaxBytesReader(w, r.Body, maxRequestBodySize)
	defer r.Body.Close()

	bodyContent, err := io.ReadAll(limitedReader)
	fileFormat := http.DetectContentType(bodyContent)
	if err != nil && !errors.Is(err, io.EOF) {
		if errors.As(err, new(*http.MaxBytesError)) {
			w.WriteHeader(http.StatusRequestEntityTooLarge)

		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			//log.Fatal(err, " ///////////////////////////")
			//errorMessage := fmt.Sprintf("Ошибка при обновлении изображения пользователя: %v", err)
			//errJson, _ := json.Marshal(map[string]string{"error": errorMessage})
			//w.Write(errJson)
		}

		return
	}

	user, err := h.usersUsecase.UpdateUserImage(r.Context(), uid, bodyContent, fileFormat)
	if err != nil {
		w.WriteHeader(http.StatusTooManyRequests)

		return
	}

	w.WriteHeader(http.StatusOK)
	userJSON, _ := json.Marshal(user)
	w.Write(userJSON)
}
