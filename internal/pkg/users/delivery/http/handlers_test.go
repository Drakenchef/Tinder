package users

import (
	mock "github.com/drakenchef/Tinder/internal/pkg/users/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
	handler := NewUsersHandler(usecase)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/list", strings.NewReader(""))
	handler.UsersList(w, r)
	assert.Equal(t, w.Code, http.StatusForbidden)
}
