package users

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestUsersList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &UsersRepo{db: db}

	expectedUsers := []models.User{
		{UID: uuid.New(), Login: "user1", Password: "pass1"},
		{UID: uuid.New(), Login: "user2", Password: "pass2"},
	}

	rows := sqlmock.NewRows([]string{"uid", "login", "passwordhash"})
	for _, user := range expectedUsers {
		rows.AddRow(user.UID, user.Login, user.Password)
	}

	mock.ExpectQuery("SELECT uid, login, passwordhash FROM users").
		WillReturnRows(rows)

	ctx := context.Background()

	users, err := r.UsersList(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(users) != len(expectedUsers) {
		t.Errorf("expected %d users, got %d", len(expectedUsers), len(users))
	}

	for i, user := range users {
		if user.UID != expectedUsers[i].UID {
			t.Errorf("unexpected UID for user %d, expected %d, got %d", i+1, expectedUsers[i].UID, user.UID)
		}

		if user.Login != expectedUsers[i].Login {
			t.Errorf("unexpected login for user %d, expected %s, got %s", i+1, expectedUsers[i].Login, user.Login)
		}

		if user.Password != expectedUsers[i].Password {
			t.Errorf("unexpected password for user %d, expected %s, got %s", i+1, expectedUsers[i].Password, user.Password)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
