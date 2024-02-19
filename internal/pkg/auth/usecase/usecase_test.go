package auth

import (
	auth "github.com/drakenchef/Tinder/internal/pkg/auth/repo"
	"reflect"
	"testing"
)

func TestNewAuthUsecase(t *testing.T) {
	type args struct {
		authRepo *auth.AuthRepo
	}
	tests := []struct {
		name string
		args args
		want *AuthUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUsecase(tt.args.authRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
