package inputPort

import "backend-food-menu-qr/core/domain"

type UserInputPort interface {
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	GetAllOwners() ([]*domain.User, error)
}
