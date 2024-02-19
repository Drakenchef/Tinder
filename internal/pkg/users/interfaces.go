package users

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/intefaces_mock.go -package=mock

type UsersUsecase interface {
	UsersList(ctx context.Context) ([]models.User, error)
}
type UsersRepo interface {
	UsersList(context.Context) ([]models.User, error)
}
