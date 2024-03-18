package users

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const basePath = "/app/images/"

type UsersRepo struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewUsersRepo(db *sql.DB, logger *zap.SugaredLogger) *UsersRepo {
	return &UsersRepo{
		db:     db,
		logger: logger,
	}
}

const ImageStoragePath = "/app/images"

func (r *UsersRepo) UsersList(ctx context.Context) ([]models.User, error) {
	utils.NameFuncLog()
	var users []models.User
	rows, err := r.db.QueryContext(ctx, "SELECT uid, login, description FROM users")
	if err != nil {
		r.logger.Info(err)
		return []models.User{}, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UID, &user.Login, &user.Description)
		if err != nil {
			r.logger.Info(err)
			return []models.User{}, err
		}
		users = append(users, user)
	}
	return users, err
}

func (r *UsersRepo) GetUser(ctx context.Context, userID uuid.UUID) (models.User, error) {
	utils.NameFuncLog()
	var user models.User
	userQuery := "SELECT uid, login, description FROM users WHERE uid = $1"
	err := r.db.QueryRowContext(ctx, userQuery, userID).Scan(&user.UID, &user.Login, &user.Description)
	if err != nil {
		r.logger.Info(err)
		return models.User{}, err
	}

	imagesQuery := "SELECT id, url, user_id FROM images WHERE user_id = $1"
	rows, err := r.db.QueryContext(ctx, imagesQuery, userID)
	if err != nil {
		r.logger.Info(err)
		return models.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var img models.Image
		if err := rows.Scan(&img.ID, &img.URL, &img.UserID); err != nil {
			r.logger.Info(err)
			return models.User{}, err
		}
		user.Images = append(user.Images, img)
	}

	return user, nil
}

func (r *UsersRepo) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	utils.NameFuncLog()
	if user.Login != "" {
		query := fmt.Sprintf("UPDATE users SET login = $1 WHERE uid = $2")
		_, err := r.db.ExecContext(ctx, query, user.Login, user.UID)
		if err != nil {
			r.logger.Info(err)
			return models.User{}, err
		}
	}
	if user.Description != "" {
		query := fmt.Sprintf("UPDATE users SET description = $1 WHERE uid = $2")
		_, err := r.db.ExecContext(ctx, query, user.Description, user.UID)
		if err != nil {
			r.logger.Info(err)
			return models.User{}, err
		}
	}
	return user, nil
}

func (r *UsersRepo) UpdateUserImage(ctx context.Context, uid uuid.UUID, imageName string) error {
	utils.NameFuncLog()
	_, err := r.db.ExecContext(ctx, "INSERT INTO images (url, user_id) VALUES ($1, $2);", imageName, uid)
	if err != nil {
		r.logger.Info(err)
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}

func (r *UsersRepo) DeleteUserImage(ctx context.Context, image string, uid uuid.UUID) error {
	utils.NameFuncLog()
	query := "DELETE FROM images WHERE url = $1 AND user_id = $2"

	_, err := r.db.ExecContext(ctx, query, image, uid)
	if err != nil {
		r.logger.Info(err)
		return err
	}

	return nil
}

func (r *UsersRepo) UpdateUserPassword(ctx context.Context, request models.ChangePassword, uid uuid.UUID) error {
	utils.NameFuncLog()
	query := fmt.Sprintf("UPDATE users SET passwordhash = $1 WHERE uid = $2")
	_, err := r.db.ExecContext(ctx, query, request.Password, uid)
	if err != nil {
		r.logger.Info(err)
		return err
	}
	return nil
}

func (r *UsersRepo) GetSaltByUid(ctx context.Context, uid uuid.UUID) (string, error) {
	utils.NameFuncLog()
	var salt string
	err := r.db.QueryRowContext(ctx, "SELECT salt FROM users WHERE uid = $1", uid).Scan(&salt)
	if err != nil {
		r.logger.Info(err)
		return "", err
	}

	return salt, nil
}

func (r *UsersRepo) DeleteUser(ctx context.Context, password models.ChangePassword, uid uuid.UUID) error {
	utils.NameFuncLog()
	query := "DELETE FROM users WHERE uid = $1 AND passwordhash = $2"

	_, err := r.db.ExecContext(ctx, query, uid, password.Password)
	if err != nil {
		r.logger.Info(err)
		return err
	}

	return nil
}
