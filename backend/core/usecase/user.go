package usecase

import (
	"backend-food-menu-qr/core/domain"
	outputPort "backend-food-menu-qr/ports/output"
	"errors"
)

type UserUseCase struct {
	userOutputPort outputPort.UserOutputPort
}

func NewUserUseCase(userOutputPort outputPort.UserOutputPort) *UserUseCase {
	return &UserUseCase{userOutputPort: userOutputPort}
}

func (u *UserUseCase) CreateUser(user *domain.User) (*domain.User, error) {
	if user.Fullname == "" || user.Email == "" || user.Password == "" || user.DateOfBirth == "" {
		return nil, errors.New("invalid data for user")
	}

	existingUser, err := u.userOutputPort.GetUserByID(user.ID)
	if err != nil {
		return nil, errors.New("database internal error")
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	return u.userOutputPort.SaveUser(user)
}

func (u *UserUseCase) UpdateUser(user *domain.User) (*domain.User, error) {
	updatedUser, err := u.userOutputPort.GetUserByID(user.ID)
	if err != nil {
		return nil, err
	}

	if updatedUser == nil {
		return nil, errors.New("user not found")
	}

	if user.Fullname != "" {
		updatedUser.Fullname = user.Fullname
	}

	if user.Email != "" {
		updatedUser.Email = user.Email
	}

	if user.DateOfBirth != "" {
		updatedUser.DateOfBirth = user.DateOfBirth
	}

	updatedUser, err = u.userOutputPort.SaveUser(updatedUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *UserUseCase) GetUserByID(id string) (*domain.User, error) {
	user, err := u.userOutputPort.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) GetAllUsers() ([]*domain.User, error) {
	users, err := u.userOutputPort.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
