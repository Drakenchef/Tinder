package likes

import (
	"bytes"
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/likes/mocks"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLikeUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikesUsecase := mock.NewMockLikesUsecase(ctrl)
	mockLogger := zap.NewExample().Sugar()

	likes := models.LikesUID{UIDFirstLike: "1111", UIDSecondLike: "2222"}
	body, _ := json.Marshal(likes)

	req, err := http.NewRequest("POST", "/likeUser", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	mockLikesUsecase.EXPECT().LikeUser(gomock.Any(), gomock.Any()).Return(nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewLikesHandler(mockLikesUsecase, mockLogger).LikeUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMutualLikeUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikesUsecase := mock.NewMockLikesUsecase(ctrl)
	mockLogger := zap.NewExample().Sugar()

	likes := models.LikesUID{UIDFirstLike: "1111", UIDSecondLike: "2222"}
	body, _ := json.Marshal(likes)

	req, err := http.NewRequest("POST", "/mutualLikeUser", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	mockLikesUsecase.EXPECT().MutualLikeUser(gomock.Any(), gomock.Any()).Return(nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewLikesHandler(mockLikesUsecase, mockLogger).MutualLikeUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
