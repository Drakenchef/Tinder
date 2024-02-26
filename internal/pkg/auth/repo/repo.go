package auth

import (
	"context"
	"database/sql"
	"errors"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/google/uuid"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(ctx context.Context, user models.SignInInput, salt string) error {
	uid := uuid.New()
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (uid, login, passwordhash, salt) VALUES ($1, $2, $3, $4)", uid, user.Login, user.Password, salt)
	if err != nil {
		return errors.New("failed to create user in database")
	}
	return nil
}
func (r *AuthRepo) GetUser(ctx context.Context, login, password string) (models.User, error) {
	var user models.User
	err := r.db.QueryRowContext(ctx, "SELECT uid, login, passwordhash FROM users WHERE login = $1 AND passwordhash = $2", login, password).Scan(&user.UID, &user.Login, &user.Password)
	if err != nil {
		return models.User{}, errors.New("failed to get user from database")
	}

	return user, nil
}
func (r *AuthRepo) GetSaltByLogin(ctx context.Context, login string) (string, error) {
	var salt string
	err := r.db.QueryRowContext(ctx, "SELECT salt FROM users WHERE login = $1", login).Scan(&salt)
	if err != nil {
		return "", errors.New("failed to get user's salt from database")
	}

	return salt, nil
}
