package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	UserRole  Role = "user"
	OwnerRole Role = "owner"
	AdminRole Role = "admin"
)

type User struct {
	ID          string `gorm:"type:uuid;primaryKey" json:"id"`
	Email       string `gorm:"unique;not null" json:"email"`
	Fullname    string `json:"fullname"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	Role        Role   `json:"role"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return nil
}
