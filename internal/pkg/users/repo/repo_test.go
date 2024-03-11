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
		{UID: uuid.New(), Login: "user1", Image: "img1", Description: "description1"},
		{UID: uuid.New(), Login: "user2", Image: "img2", Description: "description2"},
	}

	rows := sqlmock.NewRows([]string{"uid", "login", "img", "description"}).
		AddRow(expectedUsers[0].UID, expectedUsers[0].Login, expectedUsers[0].Image, expectedUsers[0].Description).
		AddRow(expectedUsers[1].UID, expectedUsers[1].Login, expectedUsers[1].Image, expectedUsers[1].Description)
	mock.ExpectQuery("SELECT uid, login, img, description FROM users").
		WillReturnRows(rows)

	ctx := context.Background()

	result, err := r.UsersList(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(result) != 2 {
		t.Errorf("unexpected number of users, expected 2, got %d", len(result))
	}

	for i, user := range result {
		if user.UID != expectedUsers[i].UID {
			t.Errorf("unexpected UID, expected %s, got %s", expectedUsers[i].UID, user.UID)
		}
		if user.Login != expectedUsers[i].Login {
			t.Errorf("unexpected login, expected %s, got %s", expectedUsers[i].Login, user.Login)
		}
		if user.Image != expectedUsers[i].Image {
			t.Errorf("unexpected image, expected %s, got %s", expectedUsers[i].Image, user.Image)
		}
		if user.Description != expectedUsers[i].Description {
			t.Errorf("unexpected description, expected %s, got %s", expectedUsers[i].Description, user.Description)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &UsersRepo{db: db}

	expectedUser := models.User{
		UID:         uuid.New(),
		Login:       "user1",
		Image:       "image1",
		Description: "description1",
	}

	rows := sqlmock.NewRows([]string{"uid", "login", "img", "description"}).
		AddRow(expectedUser.UID, expectedUser.Login, expectedUser.Image, expectedUser.Description)

	mock.ExpectQuery("SELECT uid, login, img, description FROM users").
		WithArgs(expectedUser.UID).
		WillReturnRows(rows)

	ctx := context.Background()

	user, err := r.GetUser(ctx, expectedUser.UID)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if user.UID != expectedUser.UID {
		t.Errorf("unexpected UID, expected %s, got %s", expectedUser.UID, user.UID)
	}

	if user.Login != expectedUser.Login {
		t.Errorf("unexpected login, expected %s, got %s", expectedUser.Login, user.Login)
	}

	if user.Image != expectedUser.Image {
		t.Errorf("unexpected image, expected %s, got %s", expectedUser.Image, user.Image)
	}

	if user.Description != expectedUser.Description {
		t.Errorf("unexpected description, expected %s, got %s", expectedUser.Description, user.Description)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &UsersRepo{db: db}

	user := models.User{
		UID:         uuid.New(),
		Login:       "newlogin",
		Description: "newdescription",
	}

	mock.ExpectExec("UPDATE users SET login =(.+) WHERE uid = (.+)").
		WithArgs(user.Login, user.UID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE users SET description =(.+) WHERE uid = (.+)").
		WithArgs(user.Description, user.UID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()

	updatedUser, err := r.UpdateUser(ctx, user)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if updatedUser.UID != user.UID {
		t.Errorf("unexpected UID, expected %s, got %s", user.UID, updatedUser.UID)
	}

	if updatedUser.Login != user.Login {
		t.Errorf("unexpected login, expected %s, got %s", user.Login, updatedUser.Login)
	}

	if updatedUser.Description != user.Description {
		t.Errorf("unexpected description, expected %s, got %s", user.Description, updatedUser.Description)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestUpdateUserImage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &UsersRepo{db: db}

	userID := uuid.New()
	imageName := "newimage.jpg"

	mock.ExpectExec("UPDATE users SET img=").
		WithArgs(imageName, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()

	err = r.UpdateUserImage(ctx, userID, imageName)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &UsersRepo{db: db}

	request := models.ChangePassword{
		Password: "newpassword",
	}
	userID := uuid.New()

	mock.ExpectExec("UPDATE users SET passwordhash =").
		WithArgs(request.Password, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()

	err = r.UpdateUserPassword(ctx, request, userID)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetSaltByUid(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &UsersRepo{db: db}

	userID := uuid.New()
	salt := "abc123"

	rows := sqlmock.NewRows([]string{"salt"}).AddRow(salt)
	mock.ExpectQuery("SELECT salt FROM users WHERE uid").
		WithArgs(userID).
		WillReturnRows(rows)

	ctx := context.Background()

	result, err := r.GetSaltByUid(ctx, userID)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if result != salt {
		t.Errorf("unexpected result, expected %s, got %s", salt, result)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &UsersRepo{db: db}

	password := models.ChangePassword{
		Password: "password123",
	}
	userID := uuid.New()

	mock.ExpectExec("DELETE FROM users WHERE uid =").
		WithArgs(userID, password.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()

	err = r.DeleteUser(ctx, password, userID)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
