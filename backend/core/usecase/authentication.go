package usecase

import (
	"backend-food-menu-qr/config"
	"backend-food-menu-qr/core/domain"
	outputPort "backend-food-menu-qr/ports/output"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationUseCase struct {
	userOutputPort outputPort.UserOutputPort
}

func NewAuthenticationUseCase(userOutputPort outputPort.UserOutputPort) *AuthenticationUseCase {
	return &AuthenticationUseCase{userOutputPort: userOutputPort}
}

func (c *AuthenticationUseCase) Register(user *domain.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	_, err = c.userOutputPort.SaveUser(user)
	return err
}

func (c *AuthenticationUseCase) Login(email string, password string) (*domain.User, string, error) {
	user, err := c.userOutputPort.GetUserByEmail(email)
	if err != nil {
		return &domain.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &domain.User{},"", err
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return &domain.User{},"", err
	}
	return user,token, nil
}

func GenerateJWT(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := config.AppConfig.SECRET_KEY_JWT
	secretKey := []byte(key)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
