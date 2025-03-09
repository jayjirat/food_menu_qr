package outputPort

import "backend-food-menu-qr/core/domain"

type UserOutputPort interface {
	SaveUser(user *domain.User) (*domain.User, error)
	DeleteUser(user *domain.User) error
	GetAllUsers() ([]*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
