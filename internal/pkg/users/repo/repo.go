package users

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/google/uuid"
	"log"
)

const basePath = "/app/images/"

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

const ImageStoragePath = "/app/images"

func (r *UsersRepo) UsersList(ctx context.Context) ([]models.User, error) {
	var users []models.User
	//limit 10 offset 10
	rows, err := r.db.QueryContext(ctx, "SELECT uid, login, passwordhash FROM users ")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UID, &user.Login, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, err
}

func (r *UsersRepo) GetUser(ctx context.Context, userID uuid.UUID) (models.User, error) {
	var user models.User
	query := "SELECT uid, login, img, description FROM users WHERE uid = $1"
	row := r.db.QueryRowContext(ctx, query, userID)
	err := row.Scan(&user.UID, &user.Login, &user.Image, &user.Description)
	if err != nil {
		//
	}
	log.Print(user)
	return user, nil
}

func (r *UsersRepo) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	if user.Login != "" {
		query := fmt.Sprintf("UPDATE users SET login = $1 WHERE uid = $2")
		_, err := r.db.ExecContext(ctx, query, user.Login)
		if err != nil {
			//return models.User{}, err
		}
	}
	if user.Description != "" {
		query := fmt.Sprintf("UPDATE users SET description = $1 WHERE uid = $2")
		_, err := r.db.ExecContext(ctx, query, user.Description)
		if err != nil {
			//return models.User{}, err
		}
	}
	return user, nil
}

func (r *UsersRepo) UpdateUserImage(ctx context.Context, uid uuid.UUID, imageName string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET img=$1 WHERE uid=$2;", imageName, uid)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}
