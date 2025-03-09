package inputPort

import "backend-food-menu-qr/core/domain"

type AuthenticationPort interface {
	Register(user *domain.User) error
	Login(email string, password string) (string, error)
	Logout() error
}
