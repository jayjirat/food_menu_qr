package httpAdapter

import (
	"backend-food-menu-qr/core/domain"
	inputPort "backend-food-menu-qr/ports/input"

	"github.com/gofiber/fiber/v2"
)

type AuthenticationAdapter struct {
	authenticationInputPort inputPort.AuthenticationPort
}

func NewAuthenticationAdapter(authenticationInputPort inputPort.AuthenticationPort) *AuthenticationAdapter {
	return &AuthenticationAdapter{authenticationInputPort: authenticationInputPort}
}

func (a *AuthenticationAdapter) Register(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	if err := a.authenticationInputPort.Register(&user); err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Register Successfully",
	})
}
func (a *AuthenticationAdapter) Login(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	token, err := a.authenticationInputPort.Login(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// JWT authentication
	c.Set("Authorization", "Bearer "+token)
	return c.JSON(fiber.Map{
		"message": "Logged In Successfully",
	})

}
