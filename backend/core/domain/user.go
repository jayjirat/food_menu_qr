package domain

type User struct {
	ID          uint64 `json:"id"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Password    string `json:"-"`
	DateOfBirth string `json:"date_of_birth"`
}

type UserService interface {
	GetUserByID(id int) (*User, error)
	GetAllUsers() ([]User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int) error
}

type AuthenticationService interface {
	Login(user *User) error
	Logout() error
	IsEmailExist(email string) bool
	IsPasswordValid(password, hashedPassword string) bool
}
