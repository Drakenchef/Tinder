package users

import (
	"encoding/json"
	users "github.com/drakenchef/Tinder/internal/pkg/users/usecase"
	"github.com/pkg/errors"
	"net/http"
)

type UsersHandler struct {
	usersUsecase *users.UsersUsecase
}

func NewUsersHandler(usersUsecase *users.UsersUsecase) *UsersHandler {
	return &UsersHandler{usersUsecase: usersUsecase}
}

func (h *UsersHandler) UsersList(w http.ResponseWriter, r *http.Request) {
	users, err := h.usersUsecase.UsersList()
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to sign up user").Error(), http.StatusInternalServerError)
		return
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(usersJSON)
}
