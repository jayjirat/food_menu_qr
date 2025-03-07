package repositoryAdapter

import (
	"backend-food-menu-qr/core/domain"

	"gorm.io/gorm"
)

type UserOutputAdapter struct {
	db *gorm.DB
}

func NewUserOutputAdapter(db *gorm.DB) *UserOutputAdapter {
	return &UserOutputAdapter{db: db}
}

func (u *UserOutputAdapter) SaveUser(user *domain.User) (*domain.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserOutputAdapter) DeleteUser(user *domain.User) error {
	if err := u.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserOutputAdapter) GetUserByID(id string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("id =?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserOutputAdapter) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
