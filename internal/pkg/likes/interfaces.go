package likes

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/intefaces_mock.go -package=mock
type LikesUsecase interface {
	LikeUser(context.Context, models.Likes) error
	MutualLikeUser(context.Context, models.Likes) error
}
type LikesRepo interface {
	LikeUser(context.Context, models.Likes) error
	MutualLikeUser(context.Context, models.Likes) error
}
