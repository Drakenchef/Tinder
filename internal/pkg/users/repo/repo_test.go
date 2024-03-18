package users

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"log"
	"regexp"
	"testing"
)

func TestUsersList(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mockUsers := []models.User{
		{UID: uuid.New(), Login: "user1", Description: "A user"},
		{UID: uuid.New(), Login: "user2", Description: "Another user"},
	}

	rows := sqlmock.NewRows([]string{"uid", "login", "description"}).
		AddRow(mockUsers[0].UID, mockUsers[0].Login, mockUsers[0].Description).
		AddRow(mockUsers[1].UID, mockUsers[1].Login, mockUsers[1].Description)

	mock.ExpectQuery("SELECT uid, login, description FROM users").WillReturnRows(rows)
	mockLogger := zap.NewExample().Sugar()
	repo := NewUsersRepo(db, mockLogger)
	users, err := repo.UsersList(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, len(mockUsers), len(users))
	assert.Equal(t, mockUsers, users)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userID := uuid.New()
	userRows := sqlmock.NewRows([]string{"uid", "login", "description"}).
		AddRow(userID, "testuser", "Test Description")
	imageRows := sqlmock.NewRows([]string{"id", "url", "user_id"}).
		AddRow(1, "http://example.com/image.jpg", userID)

	mock.ExpectQuery("SELECT uid, login, description FROM users WHERE uid =").
		WithArgs(userID).
		WillReturnRows(userRows)

	mock.ExpectQuery("SELECT id, url, user_id FROM images WHERE user_id =").
		WithArgs(userID).
		WillReturnRows(imageRows)

	ctx := context.Background()
	mockLogger := zap.NewExample().Sugar()
	repo := NewUsersRepo(db, mockLogger)
	user, err := repo.GetUser(ctx, userID)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if user.UID != userID {
		t.Errorf("expected userID %s, got %s", userID, user.UID)
	}
	if len(user.Images) != 1 {
		t.Errorf("expected 1 image, got %d", len(user.Images))
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
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockLogger := zap.NewExample().Sugar()
	repo := NewUsersRepo(db, mockLogger)

	testUID := uuid.New()
	testImageName := "test-image.jpg"

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO images (url, user_id) VALUES ($1, $2);")).
		WithArgs(testImageName, testUID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateUserImage(context.Background(), testUID, testImageName)
	if err != nil {
		t.Errorf("error was not expected while updating image: %s", err)
	}

	// Проверяем, выполнены ли все ожидания
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

func TestDeleteUserImage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger := zap.NewExample().Sugar()
	r := &UsersRepo{db: db, logger: logger}

	image := "http://example.com/image.jpg"
	uid := uuid.New()

	mock.ExpectExec("DELETE FROM images WHERE url = (.+) AND user_id = (.+)").
		WithArgs(image, uid).
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := context.Background()

	err = r.DeleteUserImage(ctx, image, uid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
