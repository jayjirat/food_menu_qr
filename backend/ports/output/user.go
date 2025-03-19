package outputPort

import "backend-food-menu-qr/core/domain"

type UserOutputPort interface {
	SaveUser(user *domain.User) (*domain.User, error)
	DeleteUser(userId string) error
	GetUserByUserId(userId string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	GetAllOwners() ([]*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
