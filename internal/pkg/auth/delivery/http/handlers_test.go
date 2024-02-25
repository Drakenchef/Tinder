package auth

import (
	"fmt"
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"html"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUpOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockAuthUsecase(ctrl)
	usecase.EXPECT().CreateUser(gomock.Any(), models.SignInInput{Login: "Moto", Password: "123"}).Return(nil)
	handler := NewAuthHandler(usecase)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("{\n    \"login\":\"Moto\",\n    \"password\":\"123\"\n}"))
	handler.SignUp(w, r)
	assert.Equal(t, w.Code, http.StatusCreated)
}

func TestSignUpFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockAuthUsecase(ctrl)
	t.Run("EmptyRequestBodyInSignUp", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/signup", nil)
		w := httptest.NewRecorder()
		handler := NewAuthHandler(usecase)
		handler.SignUp(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
func TestSignUpInternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockAuthUsecase(ctrl)
	usecase.EXPECT().CreateUser(gomock.Any(), models.SignInInput{Login: "Moto", Password: "123"}).Return(errors.New("internal server error"))
	handler := NewAuthHandler(usecase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("{\n    \"login\":\"Moto\",\n    \"password\":\"123\"\n}"))

	handler.SignUp(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSignInOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockAuthUsecase(ctrl)
	usecase.EXPECT().GenerateToken(gomock.Any(), models.SignInInput{Login: "Moto", Password: "123"}).Return("", nil)
	req := httptest.NewRequest(http.MethodPost, "/signin",
		strings.NewReader("{ \"login\": \"Moto\", \"password\": \"123\" }"))
	w := httptest.NewRecorder()
	handler := NewAuthHandler(usecase)
	handler.SignIn(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignInFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockAuthUsecase(ctrl)

	t.Run("EmptyRequestBodyInSignIn", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/signin", nil)
		w := httptest.NewRecorder()
		handler := NewAuthHandler(usecase)
		handler.SignIn(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
func TestSignInInternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockAuthUsecase(ctrl)
	usecase.EXPECT().GenerateToken(gomock.Any(), models.SignInInput{Login: "Moto", Password: "123"}).Return("", errors.New("internal server error"))

	req := httptest.NewRequest(http.MethodPost, "/signin",
		strings.NewReader("{ \"login\": \"Moto\", \"password\": \"123\" }"))
	w := httptest.NewRecorder()
	handler := NewAuthHandler(usecase)
	handler.SignIn(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCheckAuthAuthorized(t *testing.T) {
	t.Skip("work in progress")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockAuthUsecase(ctrl)
	handler := NewAuthHandler(usecase)

	req := httptest.NewRequest(http.MethodGet, "/checkauth", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDg2NjA3NjIsImlhdCI6MTcwODYxNzU2MiwidWlkIjoiN2FhMTMxNDEtODAwMy00Zjc5LWI4ZTUtYThmMGM5MDIwYmJjIn0.Z5PjZjRxDo9IECMxPsLPiqpeQ5kT0KTz9lMGGqPTyfE")
	w := httptest.NewRecorder()

	handler.CheckAuth(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuthUnauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockAuthUsecase(ctrl)
	handler := NewAuthHandler(usecase)

	req := httptest.NewRequest(http.MethodGet, "/checkauth", nil)
	w := httptest.NewRecorder()

	handler.CheckAuth(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestFoo(t *testing.T) {
	input := "<>&\"'"
	output := html.EscapeString(input)
	fmt.Println(output)
}
