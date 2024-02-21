package auth

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockAuthRepo(ctrl)
	usecase := NewAuthUsecase(repo)
	ctx := context.Background()
	emptyLoginUser := models.SignInInput{
		Login:    "",
		Password: "123",
	}
	err := usecase.CreateUser(ctx, emptyLoginUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Login and password are required")
	emptyPasswordUser := models.SignInInput{
		Login:    "Moto",
		Password: "",
	}
	err = usecase.CreateUser(ctx, emptyPasswordUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Login and password are required")
	goodPassNLogin := models.SignInInput{
		Login:    "Moto",
		Password: "123",
	}
	repo.EXPECT().CreateUser(gomock.Any(), goodPassNLogin, gomock.Any()).Return(nil)
	err = usecase.CreateUser(ctx, goodPassNLogin)
	assert.NoError(t, err)
}

func TestGenerateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockAuthRepo(ctrl)
	usecase := NewAuthUsecase(repo)
	ctx := context.Background()
	input := models.SignInInput{
		Login:    "Moto",
		Password: "123",
	}
	salt := "salt"
	repo.EXPECT().GetSaltByLogin(ctx, input.Login).Return(salt, nil)
	repo.EXPECT().GetUser(ctx, input.Login, gomock.Any()).Return(models.User{}, nil)
	tokenStr, err := usecase.GenerateToken(ctx, input)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

}
