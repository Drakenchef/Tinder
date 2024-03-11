package users

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/intefaces_mock.go -package=mock

type UsersUsecase interface {
	UsersList(context.Context) ([]models.User, error)
	GetUser(context.Context, uuid.UUID) (models.User, error)
	UpdateUser(context.Context, models.User) (models.User, error)
	UpdateUserImage(context.Context, uuid.UUID, []byte, string) (models.User, error)
	UpdateUserPassword(context.Context, models.ChangePassword, uuid.UUID) error
	DeleteUser(context.Context, models.ChangePassword, uuid.UUID) error
}
type UsersRepo interface {
	UsersList(context.Context) ([]models.User, error)
	GetUser(context.Context, uuid.UUID) (models.User, error)
	UpdateUser(context.Context, models.User) (models.User, error)
	UpdateUserImage(context.Context, uuid.UUID, string) error
	GetSaltByUid(context.Context, uuid.UUID) (string, error)
	UpdateUserPassword(context.Context, models.ChangePassword, uuid.UUID) error
	DeleteUser(context.Context, models.ChangePassword, uuid.UUID) error
}
