package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/auth"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"time"
)

const (
	signinkey = "qweeqwrqw2131e2r3t22trtrfg42"
	tokenTTL  = 12 * time.Hour
)

type AuthUsecase struct {
	authRepo auth.AuthRepo
	auther   auth.Auther
}

func NewAuthUsecase(authRepo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{authRepo: authRepo, auther: &Auther{}}
}

type tokenClaims struct {
	jwt.StandardClaims
	UID uuid.UUID `json:"uid"`
}

func (u *AuthUsecase) CreateUser(ctx context.Context, user models.SignInInput) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return errors.New("validation failed")
	}
	p := bluemonday.UGCPolicy()
	user.Login = p.Sanitize(user.Login)
	user.Password = p.Sanitize(user.Password)
	salt := utils.GenerateRandomSalt()
	user.Password = u.auther.GeneratePasswordHash(user.Password + salt)
	err := u.authRepo.CreateUser(ctx, user, salt)
	if err != nil {
		return errors.New("failed to create user in repository")
	}

	return nil
}

func (u *AuthUsecase) GenerateToken(ctx context.Context, input models.SignInInput) (string, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return "", errors.New("validation failed")
	}

	p := bluemonday.UGCPolicy()
	input.Login = p.Sanitize(input.Login)
	input.Password = p.Sanitize(input.Password)

	salt, err := u.authRepo.GetSaltByLogin(ctx, input.Login)
	if err != nil {
		return "", err
	}
	user, err := u.authRepo.GetUser(ctx, input.Login, u.auther.GeneratePasswordHash(input.Password+salt))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UID,
	})
	return token.SignedString([]byte(signinkey))
}
