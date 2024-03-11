package likes

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/google/uuid"
	"log"
	"testing"
	"time"
)

func TestLikeUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &LikesRepo{db: db}

	likes := models.Likes{UIDFirstLike: uuid.New(), UIDSecondLike: uuid.New(), Date: time.Now(), Mutual: false}

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(false)
	mock.ExpectQuery("SELECT EXISTS").WithArgs(likes.UIDFirstLike, likes.UIDSecondLike).WillReturnRows(rows)

	mock.ExpectExec("INSERT INTO likes").WithArgs(likes.UIDFirstLike, likes.UIDSecondLike, likes.Date, likes.Mutual).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()

	err = r.LikeUser(ctx, likes)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMutualLikeUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := &LikesRepo{db: db}

	likes := models.Likes{UIDFirstLike: uuid.New(), UIDSecondLike: uuid.New()}

	mock.ExpectExec("UPDATE likes SET mutual = true").WithArgs(likes.UIDFirstLike, likes.UIDSecondLike, likes.UIDSecondLike, likes.UIDFirstLike).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()

	err = r.MutualLikeUser(ctx, likes)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
