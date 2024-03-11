package likes

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/likes"
	"github.com/drakenchef/Tinder/internal/utils"
	"go.uber.org/zap"
)

type LikesUsecase struct {
	likesRepo likes.LikesRepo
	logger    *zap.SugaredLogger
}

func NewLikesUsecase(likesRepo likes.LikesRepo, logger *zap.SugaredLogger) *LikesUsecase {
	return &LikesUsecase{
		likesRepo: likesRepo,
		logger:    logger,
	}
}

func (u *LikesUsecase) LikeUser(ctx context.Context, likes models.Likes) error {
	utils.NameFuncLog()
	return u.likesRepo.LikeUser(ctx, likes)
}

func (u *LikesUsecase) MutualLikeUser(ctx context.Context, likes models.Likes) error {
	utils.NameFuncLog()
	return u.likesRepo.MutualLikeUser(ctx, likes)
}
