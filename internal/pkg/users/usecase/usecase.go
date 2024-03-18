package users

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/users"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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
	logger    *zap.SugaredLogger
}

func NewUsersUsecase(usersRepo users.UsersRepo, logger *zap.SugaredLogger) *UsersUsecase {
	return &UsersUsecase{
		usersRepo: usersRepo,
		logger:    logger,
	}
}

func (u *UsersUsecase) UsersList(ctx context.Context) ([]models.User, error) {
	utils.NameFuncLog()
	users, err := u.usersRepo.UsersList(ctx)
	if err != nil {
		return nil, errors.New("failed to find users in repository")
	}
	return users, nil
}

func (u *UsersUsecase) GetUser(ctx context.Context, userID uuid.UUID) (models.User, error) {
	utils.NameFuncLog()
	return u.usersRepo.GetUser(ctx, userID)
}

func (u *UsersUsecase) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	utils.NameFuncLog()
	return u.usersRepo.UpdateUser(ctx, user)
}

func (u *UsersUsecase) UpdateUserImage(ctx context.Context, uid uuid.UUID, filePhotoByte []byte, fileType string) (models.User, error) {
	utils.NameFuncLog()
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
	imageUID := uuid.New()
	imageName := imageUID.String() + "." + fileExtension

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
func (u *UsersUsecase) DeleteUserImage(ctx context.Context, image string, uid uuid.UUID) error {
	utils.NameFuncLog()
	err := u.usersRepo.DeleteUserImage(ctx, image, uid)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersUsecase) UpdateUserPassword(ctx context.Context, request models.ChangePassword, uid uuid.UUID) error {
	utils.NameFuncLog()
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return errors.New("validation failed")
	}
	p := bluemonday.UGCPolicy()
	request.Password = p.Sanitize(request.Password)

	salt, err := u.usersRepo.GetSaltByUid(ctx, uid)
	if err != nil {
		return err
	}

	hasher := md5.New()
	hasher.Write([]byte(request.Password + salt))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	request.Password = hashedPassword
	err = u.usersRepo.UpdateUserPassword(ctx, request, uid)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersUsecase) DeleteUser(ctx context.Context, password models.ChangePassword, uid uuid.UUID) error {
	utils.NameFuncLog()
	salt, err := u.usersRepo.GetSaltByUid(ctx, uid)
	if err != nil {
		return err
	}
	hasher := md5.New()
	hasher.Write([]byte(password.Password + salt))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	password.Password = hashedPassword

	err = u.usersRepo.DeleteUser(ctx, password, uid)
	if err != nil {
		return err
	}
	return nil
}
