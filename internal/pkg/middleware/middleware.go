package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

const (
	signinkey = "qweeqwrqw2131e2r3t22trtrfg42"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UID uuid.UUID `json:"uid"`
}

func CheckAuth(r *http.Request) (uuid.UUID, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return uuid.Nil, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return uuid.Nil, errors.New("invalid auth header")
	}

	if headerParts[0] != "Bearer" {
		return uuid.Nil, errors.New("incorrect auth header")
	}

	if headerParts[1] == "" {
		return uuid.Nil, errors.New("token is empty")
	}
	userId, err := ParseToken(headerParts[1])
	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}

func ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signinkey), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.Nil, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UID, nil
}
