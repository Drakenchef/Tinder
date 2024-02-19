package auth

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/intefaces_mock.go -package=mock

type AuthUsecase interface {
	CreateUser(context.Context, models.SignInInput) error
	GenerateToken(context.Context, string, string) (string, error)
}
type AuthRepo interface {
	CreateUser(context.Context, models.SignInInput, string) error
	GetUser(context.Context, string, string) (models.User, error)
	GetSaltByLogin(context.Context, string) (string, error)
}
