package auth

import (
	"bytes"
	"encoding/json"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase := mock.NewMockAuthUsecase(ctrl)
	mockLogger := zap.NewExample().Sugar()

	user := models.SignInInput{Login: "test", Password: "password"}
	body, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	mockAuthUsecase.EXPECT().CreateUser(gomock.Any(), gomock.Eq(user)).Return(nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewAuthHandler(mockAuthUsecase, mockLogger).SignUp)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase := mock.NewMockAuthUsecase(ctrl)
	mockLogger := zap.NewExample().Sugar()

	user := models.SignInInput{Login: "testvalid", Password: "passwordvalid"}
	body, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/signin", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	expectedToken := "some-test-token"
	mockAuthUsecase.EXPECT().GenerateToken(gomock.Any(), user).Return(expectedToken, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewAuthHandler(mockAuthUsecase, mockLogger).SignIn)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	cookieFound := false
	for _, cookie := range rr.Result().Cookies() {
		if cookie.Name == "token" && cookie.Value != "" {
			cookieFound = true
		}
	}
	if !cookieFound {
		t.Fatal("auth_token cookie not set in response")
	}

}
