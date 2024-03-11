package likes

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/utils"
	"go.uber.org/zap"
)

const basePath = "/app/images/"

type LikesRepo struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewLikesRepo(db *sql.DB, logger *zap.SugaredLogger) *LikesRepo {
	return &LikesRepo{
		db:     db,
		logger: logger,
	}
}

func (r *LikesRepo) LikeUser(ctx context.Context, likes models.Likes) error {
	utils.NameFuncLog()
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM likes WHERE uidfirstlike = $1 AND uidsecondlike = $2)", likes.UIDFirstLike, likes.UIDSecondLike).Scan(&exists)
	if err != nil {
		r.logger.Info(err)
		return err
	}
	if exists {
		return fmt.Errorf("like between uidfirst=%d and uidsecond=%d already exists", likes.UIDFirstLike, likes.UIDSecondLike)
	} else {
		_, err = r.db.ExecContext(ctx, "INSERT INTO likes (uidfirstlike, uidsecondlike, Date, Mutual) VALUES ($1, $2, $3, $4)", likes.UIDFirstLike, likes.UIDSecondLike, likes.Date, false)
		if err != nil {
			r.logger.Info(err)
			return err
		}
		return nil
	}
}

func (r *LikesRepo) MutualLikeUser(ctx context.Context, likes models.Likes) error {
	utils.NameFuncLog()
	_, err := r.db.ExecContext(ctx, "UPDATE likes SET mutual = true WHERE (uidfirstlike = $1 AND uidsecondlike = $2) OR (uidfirstlike = $3 AND uidsecondlike = $4)", likes.UIDFirstLike, likes.UIDSecondLike, likes.UIDSecondLike, likes.UIDFirstLike)
	if err != nil {
		r.logger.Info(err)
		return err
	}
	return nil
}
