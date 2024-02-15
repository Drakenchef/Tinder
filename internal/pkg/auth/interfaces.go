package auth

type AuthUsecase interface {
	Login(username, password string) (string, error)
	Logout(token string) error
}
