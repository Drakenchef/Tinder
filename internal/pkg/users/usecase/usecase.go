package users

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/users"
	"github.com/pkg/errors"
)

type UsersUsecase struct {
	usersRepo users.UsersRepo
}

func NewUsersUsecase(usersRepo users.UsersRepo) *UsersUsecase {
	return &UsersUsecase{usersRepo: usersRepo}
}

func (u *UsersUsecase) UsersList(ctx context.Context) ([]models.User, error) {
	users, err := u.usersRepo.UsersList(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find users in repository")
	}
	return users, nil
}
