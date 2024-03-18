package users

import (
	"bytes"
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/users/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUsersList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockUsersUsecase(ctrl)
	usecase.EXPECT().UsersList(gomock.Any()).Return(nil, nil).AnyTimes()
	var logger *zap.SugaredLogger
	handler := NewUsersHandler(usecase, logger)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/list", strings.NewReader(""))
	handler.UsersList(w, r)
	assert.Equal(t, w.Code, http.StatusOK)
}
func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// создание мока usecase
	usecase := mock.NewMockUsersUsecase(ctrl)
	// создание заглушки логгера вместо nil
	logger := zap.NewExample().Sugar()

	handler := NewUsersHandler(usecase, logger)

	testUserId := uuid.New()
	testUser := models.User{
		UID:   testUserId,
		Login: "TestUserLogin",
	}
	usecase.EXPECT().GetUser(gomock.Any(), gomock.Eq(testUserId)).Return(testUser, nil).Times(1)

	reqBody, _ := json.Marshal(map[string]string{
		"uid": testUserId.String(),
	})
	r := httptest.NewRequest(http.MethodGet, "/get", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()
	r.Header.Set("uid", testUserId.String())

	handler.GetUser(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUsersUsecase(ctrl)
	handler := NewUsersHandler(usecase, nil)

	updateUser := models.User{
		UID:   uuid.New(),
		Login: "Updated Test User",
	}

	usecase.EXPECT().UpdateUser(gomock.Any(), gomock.Eq(updateUser)).Return(updateUser, nil).Times(1)

	reqBody, _ := json.Marshal(updateUser)
	r := httptest.NewRequest(http.MethodPut, "/update", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	handler.UpdateUser(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpdateUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUsersUsecase(ctrl)
	mockLogger := zap.NewExample().Sugar()
	handler := NewUsersHandler(usecase, mockLogger)

	testUserID := uuid.New()
	testRequest := models.ChangePassword{
		Password: "oldPass",
	}

	reqBody, _ := json.Marshal(testRequest)
	r := httptest.NewRequest(http.MethodPost, "/update-password", bytes.NewReader(reqBody))
	r.Header.Set("uid", testUserID.String())
	w := httptest.NewRecorder()

	usecase.EXPECT().UpdateUserPassword(gomock.Any(), testRequest, testUserID).Return(nil).Times(1)

	handler.UpdateUserPassword(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUsersUsecase(ctrl)
	mockLogger := zap.NewExample().Sugar()
	handler := NewUsersHandler(usecase, mockLogger)

	testUserID := uuid.New()
	testRequest := models.ChangePassword{
		Password: "password",
	}

	reqBody, _ := json.Marshal(testRequest)
	r := httptest.NewRequest(http.MethodPost, "/delete-user", bytes.NewReader(reqBody))
	r.Header.Set("uid", testUserID.String())
	w := httptest.NewRecorder()

	usecase.EXPECT().DeleteUser(gomock.Any(), testRequest, testUserID).Return(nil).Times(1)

	handler.DeleteUser(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUserImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUsersUsecase(ctrl)
	mockLogger := zap.NewExample().Sugar()
	handler := NewUsersHandler(usecase, mockLogger)

	testUserID := uuid.New()
	testImage := []byte("test image")
	var user models.User
	reqBody, _ := json.Marshal("")
	r := httptest.NewRequest(http.MethodPost, "/image", bytes.NewReader(reqBody))
	r.Header.Set("uid", testUserID.String())
	r.Body = ioutil.NopCloser(bytes.NewReader(testImage))
	w := httptest.NewRecorder()

	usecase.EXPECT().UpdateUserImage(gomock.Any(), testUserID, testImage, gomock.Any()).Return(user, nil).Times(1)

	handler.UpdateUserImage(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUserImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUsersUsecase(ctrl)
	mockLogger := zap.NewExample().Sugar()
	handler := NewUsersHandler(usecase, mockLogger)

	testUserID := uuid.New()
	testRequest := struct {
		URL string `json:"url"`
	}{
		URL: "http://example.com/image.jpg",
	}

	reqBody, _ := json.Marshal(testRequest)
	r := httptest.NewRequest(http.MethodPost, "/delete-user-image", bytes.NewReader(reqBody))
	r.Header.Set("uid", testUserID.String())
	w := httptest.NewRecorder()

	usecase.EXPECT().DeleteUserImage(gomock.Any(), testRequest.URL, testUserID).Return(nil).Times(1)

	handler.DeleteUserImage(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
