package auth

import (
	"database/sql"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/pkg/errors"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(user models.User) error {
	_, err := r.db.Exec("INSERT INTO users (login, passwordhash) VALUES ($1, $2)", user.Login, user.Password)
	if err != nil {
		return errors.Wrap(err, "failed to execute SQL query")
	}

	return nil
}
func (r *AuthRepo) GetUser(login, password string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT uid, login, passwordhash FROM users WHERE login = $1 AND passwordhash = $2", login, password).Scan(&user.UID, &user.Login, &user.Password)
	if err != nil {
		return models.User{}, errors.Wrap(err, "failed to get user from database")
	}

	return user, nil
}
