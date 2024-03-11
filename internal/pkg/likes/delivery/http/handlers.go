package likes

import (
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/models"
	"github.com/drakenchef/Tinder/internal/pkg/likes"
	"github.com/drakenchef/Tinder/internal/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type LikesHandler struct {
	likesUsecase likes.LikesUsecase
	logger       *zap.SugaredLogger
}

func NewLikesHandler(likesUsecase likes.LikesUsecase, logger *zap.SugaredLogger) *LikesHandler {
	return &LikesHandler{
		likesUsecase: likesUsecase,
		logger:       logger,
	}
}

func (h *LikesHandler) LikeUser(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	var input models.LikesUID
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uidFirst, _ := uuid.Parse(input.UIDFirstLike)
	uidSecond, _ := uuid.Parse(input.UIDSecondLike)
	var likes models.Likes
	likes.UIDFirstLike = uidFirst
	likes.UIDSecondLike = uidSecond
	likes.Date = time.Now()
	err := h.likesUsecase.LikeUser(r.Context(), likes)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *LikesHandler) MutualLikeUser(w http.ResponseWriter, r *http.Request) {
	utils.NameFuncLog()
	var input models.LikesUID
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uidFirst, _ := uuid.Parse(input.UIDFirstLike)
	uidSecond, _ := uuid.Parse(input.UIDSecondLike)
	var likes models.Likes
	likes.UIDFirstLike = uidFirst
	likes.UIDSecondLike = uidSecond
	likes.Date = time.Now()
	err := h.likesUsecase.MutualLikeUser(r.Context(), likes)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
