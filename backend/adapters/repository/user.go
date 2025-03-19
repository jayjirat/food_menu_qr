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

	var existingUser domain.User

	if err := u.db.Where("id=?", user.ID).First(&existingUser).Error; err == nil {
		if err := u.db.Model(&existingUser).Updates(user).Error; err != nil {
			return nil, err
		}
		return &existingUser, nil
	}

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

	if err := u.db.Where("role = ?", "user").Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserOutputAdapter) GetAllOwners() ([]*domain.User, error) {
	var users []*domain.User

	if err := u.db.Where("role = ?", "owner").Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserOutputAdapter) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("email =?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
