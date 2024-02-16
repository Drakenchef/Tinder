package users

import (
	"github.com/drakenchef/Tinder/internal/models"
	users "github.com/drakenchef/Tinder/internal/pkg/users/repo"
	"github.com/pkg/errors"
)

type UsersUsecase struct {
	usersRepo *users.UsersRepo
}

func NewUsersUsecase(usersRepo *users.UsersRepo) *UsersUsecase {
	return &UsersUsecase{usersRepo: usersRepo}
}

func (u *UsersUsecase) UsersList() ([]models.User, error) {
	users, err := u.usersRepo.UsersList()
	if err != nil {
		return nil, errors.Wrap(err, "failed to find users in repository")
	}
	return users, nil
}
