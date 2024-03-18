package users

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/users/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestUsersList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUsersRepo(ctrl)
	usecase := NewUsersUsecase(repo, &zap.SugaredLogger{})

	ctx := context.Background()

	mockUsers := []models.User{
		{UID: uuid.New(), Login: "John"},
		{UID: uuid.New(), Login: "Jane"},
		{UID: uuid.New(), Login: "Alice"},
	}

	repo.EXPECT().UsersList(gomock.Any()).Return(mockUsers, nil)

	users, err := usecase.UsersList(ctx)

	assert.NoError(t, err)
	assert.Equal(t, mockUsers, users)
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUsersRepo(ctrl)
	usecase := NewUsersUsecase(repo, &zap.SugaredLogger{})

	ctx := context.Background()
	uid := uuid.New()
	user := models.User{UID: uuid.New(), Login: "John"}
	repo.EXPECT().GetUser(gomock.Any(), uid).Return(user, nil)

	users, err := usecase.GetUser(ctx, uid)

	assert.NoError(t, err)
	assert.Equal(t, user, users)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUsersRepo(ctrl)
	usecase := NewUsersUsecase(repo, &zap.SugaredLogger{})

	ctx := context.Background()
	user := models.User{UID: uuid.New(), Login: "John", Description: " "}
	expectedUser := models.User{UID: uuid.New(), Login: "John", Description: "mydesc"}
	repo.EXPECT().UpdateUser(gomock.Any(), user).Return(expectedUser, nil)

	users, err := usecase.UpdateUser(ctx, user)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, users)
}
func TestUpdateUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUsersRepo(ctrl)
	usecase := NewUsersUsecase(repo, &zap.SugaredLogger{})

	ctx := context.Background()
	uid := uuid.New()
	salt := "salt123"
	hashedPassword := "Password123"

	repo.EXPECT().GetSaltByUid(gomock.Any(), uid).Return(salt, nil)

	repo.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Any(), uid).Do(func(ctx context.Context, req models.ChangePassword, uid uuid.UUID) {
		if req.Password != "0d8ab643cf9782a329488da96784d32a" {
			t.Errorf("Expected hashed password %s, got %s", hashedPassword, req.Password)
		}
	}).Return(nil)

	err := usecase.UpdateUserPassword(ctx, models.ChangePassword{Password: hashedPassword}, uid)

	assert.NoError(t, err)
}
func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUsersRepo(ctrl)
	usecase := NewUsersUsecase(repo, &zap.SugaredLogger{})

	ctx := context.Background()
	uid := uuid.New()
	password := models.ChangePassword{Password: "password123"}
	salt := "salt123"

	repo.EXPECT().GetSaltByUid(gomock.Any(), uid).Return(salt, nil)
	repo.EXPECT().DeleteUser(gomock.Any(), models.ChangePassword{Password: "69cf04164801421672ee35c6adbeafad"}, uid).Return(nil)

	err := usecase.DeleteUser(ctx, password, uid)

	assert.NoError(t, err)
}

func TestDeleteUserImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUsersRepo(ctrl)
	usecase := NewUsersUsecase(repo, zap.NewExample().Sugar())

	ctx := context.Background()
	uid := uuid.New()
	image := "http://example.com/image.jpg"

	repo.EXPECT().DeleteUserImage(ctx, image, uid).Return(nil)

	err := usecase.DeleteUserImage(ctx, image, uid)

	assert.NoError(t, err)
}
