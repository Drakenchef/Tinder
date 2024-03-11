package auth

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockAuthRepo(ctrl)
	mockLogger := zap.NewExample().Sugar()
	usecase := NewAuthUsecase(repo, mockLogger)
	ctx := context.Background()
	invalidLoginUser := models.SignInInput{
		Login:    "Невалидный Логин",
		Password: "ValidPassword",
	}
	err := usecase.CreateUser(ctx, invalidLoginUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
	invalidPasswordUser := models.SignInInput{
		Login:    "ValidLogin",
		Password: "Невалидный пароль",
	}
	err = usecase.CreateUser(ctx, invalidPasswordUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
	goodPassAndLogin := models.SignInInput{
		Login:    "ValidLogin",
		Password: "ValidPassword",
	}
	repo.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err = usecase.CreateUser(ctx, goodPassAndLogin)
	assert.NoError(t, err)
}

func TestGenerateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockAuthRepo(ctrl)
	usecase := NewAuthUsecase(repo, &zap.SugaredLogger{})
	ctx := context.Background()

	input := models.SignInInput{
		Login:    "ValidLogin",
		Password: "ValidPassword",
	}
	salt := "salt"
	user := models.User{UID: uuid.New()}
	repo.EXPECT().GetSaltByLogin(ctx, input.Login).Return(salt, nil)
	repo.EXPECT().GetUser(ctx, input.Login, gomock.Any()).Return(user, nil)

	tokenStr, err := usecase.GenerateToken(ctx, input)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)
}
