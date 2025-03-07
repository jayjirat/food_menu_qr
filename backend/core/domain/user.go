package domain

type User struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Email       string `gorm:"unique;not null" json:"email"`
	Fullname    string `json:"fullname"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
}
