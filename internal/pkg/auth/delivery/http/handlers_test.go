package auth

import (
	"github.com/drakenchef/Tinder/internal/models"
	mock "github.com/drakenchef/Tinder/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usecase := mock.NewMockAuthUsecase(ctrl)
	usecase.EXPECT().CreateUser(gomock.Any(), models.User{Login: "Moto", Password: "123"}).Return(nil)
	handler := NewAuthHandler(usecase)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("{\n    \"login\":\"Moto\",\n    \"password\":\"123\"\n}"))
	handler.SignUp(w, r)
	assert.Equal(t, w.Code, http.StatusCreated)
}
