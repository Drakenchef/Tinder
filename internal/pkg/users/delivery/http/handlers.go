package users

import (
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/pkg/users"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/pkg/errors"
	"net/http"
)

type UsersHandler struct {
	usersUsecase users.UsersUsecase
}

func NewUsersHandler(usersUsecase users.UsersUsecase) *UsersHandler {
	return &UsersHandler{usersUsecase: usersUsecase}
}

func (h *UsersHandler) UsersList(w http.ResponseWriter, r *http.Request) {
	_, err := utils.CheckAuth(r)
	if err == nil {
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
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
