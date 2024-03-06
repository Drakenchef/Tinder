package users

import (
	"context"
	"fmt"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/users"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

var (
	acceptingFileTypes = []string{"image/webp", "image/png", "image/jpeg"}
)

const imagePath = "/app/images/"

type UsersUsecase struct {
	usersRepo users.UsersRepo
}

func NewUsersUsecase(usersRepo users.UsersRepo) *UsersUsecase {
	return &UsersUsecase{usersRepo: usersRepo}
}

func (u *UsersUsecase) UsersList(ctx context.Context) ([]models.User, error) {
	users, err := u.usersRepo.UsersList(ctx)
	if err != nil {
		return nil, errors.New("failed to find users in repository")
	}
	return users, nil
}

func (u *UsersUsecase) GetUser(ctx context.Context, userID uuid.UUID) (models.User, error) {
	return u.usersRepo.GetUser(ctx, userID)
}

func (u *UsersUsecase) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return u.usersRepo.UpdateUser(ctx, user)
}

func (u *UsersUsecase) UpdateUserImage(ctx context.Context, uid uuid.UUID, filePhotoByte []byte, fileType string) (models.User, error) {
	if !slices.Contains(acceptingFileTypes, fileType) {
		return models.User{}, errors.New("Forbidden extension")
	}

	_, err := os.Stat(imagePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("photos upload file path [%s] is not exist, error: %w", imagePath, err)

			return models.User{}, err
		}

		return models.User{}, fmt.Errorf("bad photos file path [%s], error: %w", imagePath, err)
	}

	fileExtension := strings.TrimPrefix(fileType, "image/")

	imageName := uid.String() + "." + fileExtension

	file, err := os.Create(imagePath + imageName)
	if err != nil {
		err = fmt.Errorf("error happened in create file: %w", err)

		return models.User{}, err
	}
	defer file.Close()

	_, err = file.Write(filePhotoByte)
	if err != nil {
		err = fmt.Errorf("error happened in write to file: %w", err)

		return models.User{}, err
	}

	err = u.usersRepo.UpdateUserImage(ctx, uid, imageName)
	if err != nil {
		err = fmt.Errorf("error happened in repoUser.UpdatePhoto: %w", err)

		return models.User{}, err
	}

	profileInfo, err := u.usersRepo.GetUser(ctx, uid)

	return profileInfo, nil
}
