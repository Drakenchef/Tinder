package auth

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	mockLogger := zap.NewExample().Sugar()
	r := NewAuthRepo(db, mockLogger)

	user := models.SignInInput{
		Login:    "testUser",
		Password: "testPass",
	}
	ctx := context.Background()
	salt := "salt"
	mock.ExpectExec("INSERT INTO users").WithArgs(sqlmock.AnyArg(), user.Login, user.Password, salt, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.CreateUser(ctx, user, salt)
	if err != nil {
		t.Errorf("error was expected: %s", err)
	}
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &AuthRepo{db: db}

	expectedUser := models.User{
		UID:      uuid.New(),
		Login:    "user1",
		Password: "pass1",
	}

	rows := sqlmock.NewRows([]string{"uid", "login", "passwordhash"}).
		AddRow(expectedUser.UID, expectedUser.Login, expectedUser.Password)

	mock.ExpectQuery("SELECT uid, login, passwordhash FROM users WHERE login = (.+) AND passwordhash = (.+)").
		WithArgs(expectedUser.Login, expectedUser.Password).
		WillReturnRows(rows)

	ctx := context.Background()

	_, err = r.GetUser(ctx, expectedUser.Login, expectedUser.Password)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}

func TestGetSaltByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthRepo(db, &zap.SugaredLogger{})

	ctx := context.Background()
	login := "testUser"

	rows := sqlmock.NewRows([]string{"salt"}).
		AddRow("saltValue")
	mock.ExpectQuery("SELECT salt FROM users WHERE login = (.+)").
		WithArgs(login).
		WillReturnRows(rows)

	_, err = r.GetSaltByLogin(ctx, login)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
