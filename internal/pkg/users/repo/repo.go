package users

import (
	"context"
	"database/sql"
	"github.com/drakenchef/Tinder/internal/models"
	"log"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

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
