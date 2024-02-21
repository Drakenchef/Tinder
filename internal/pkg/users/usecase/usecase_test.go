package users

import (
	"context"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/users/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsersList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUsersRepo(ctrl)
	usecase := NewUsersUsecase(repo)

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
