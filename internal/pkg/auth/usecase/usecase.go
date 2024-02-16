package auth

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/drakenchef/Tinder/internal/models"
	auth "github.com/drakenchef/Tinder/internal/pkg/auth/repo"
	"github.com/pkg/errors"
	"time"
)

const (
	salt      = "justasaltsimbolsfordecode"
	signinkey = "qweeqwrqw2131e2r3t22trtrfg42"
	tokenTTL  = 12 * time.Hour
)

type AuthUsecase struct {
	authRepo *auth.AuthRepo
}

type tokenClaims struct {
	jwt.StandardClaims
	UID int `json:"uid"`
}

func NewAuthUsecase(authRepo *auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{authRepo: authRepo}
}

func (u *AuthUsecase) CreateUser(user models.User) error {
	if user.Login == "" || user.Password == "" {
		return errors.New("Login and password are required")
	}
	user.Password = generatePasswordHash(user.Password)
	err := u.authRepo.CreateUser(user)
	if err != nil {
		return errors.Wrap(err, "failed to create user in repository")
	}

	return nil
}

func (u *AuthUsecase) GenerateToken(login, password string) (string, error) {
	user, err := u.authRepo.GetUser(login, generatePasswordHash(password))
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

func (u *AuthUsecase) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signinkey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UID, nil
}

func generatePasswordHash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	return hashedPassword
}
