package likes

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/likes/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestLikeUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikesRepo := mock.NewMockLikesRepo(ctrl)
	mockLikesUsecase := NewLikesUsecase(mockLikesRepo, &zap.SugaredLogger{})

	ctx := context.Background()
	likes := models.Likes{UIDFirstLike: uuid.New(), UIDSecondLike: uuid.New()}
	mockLikesRepo.EXPECT().LikeUser(ctx, likes).Return(nil)

	err := mockLikesUsecase.LikeUser(ctx, likes)

	assert.NoError(t, err)
}

func TestMutualLikeUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikesRepo := mock.NewMockLikesRepo(ctrl)
	mockLikesUsecase := NewLikesUsecase(mockLikesRepo, &zap.SugaredLogger{})

	ctx := context.Background()
	likes := models.Likes{UIDFirstLike: uuid.New(), UIDSecondLike: uuid.New()}
	mockLikesRepo.EXPECT().MutualLikeUser(ctx, likes).Return(nil)

	err := mockLikesUsecase.MutualLikeUser(ctx, likes)

	assert.NoError(t, err)
}
