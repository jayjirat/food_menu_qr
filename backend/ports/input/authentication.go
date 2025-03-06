package input

import "backend-food-menu-qr/core/domain"

type AuthenticationPort interface {
	Login(user *domain.User) error
	Logout() error
	IsEmailExist(email string) bool
}
