package users

import (
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

func (r *UsersRepo) UsersList() ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query("SELECT * FROM users")
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
