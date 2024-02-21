package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/auth"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

const (
	salt      = "justasaltsimbolsfordecode"
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
	salt := utils.GenerateRandomSalt()
	if user.Login == "" || user.Password == "" {
		return errors.New("Login and password are required")
	}
	user.Password = u.auther.GeneratePasswordHash(user.Password + salt)
	err := u.authRepo.CreateUser(ctx, user, salt)
	if err != nil {
		return errors.Wrap(err, "failed to create user in repository")
	}

	return nil
}

//salt -> bd
//jwt -> cookie

func (u *AuthUsecase) GenerateToken(ctx context.Context, input models.SignInInput) (string, error) {
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
