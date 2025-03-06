package http

import (
	"backend-food-menu-qr/core/domain"
	"backend-food-menu-qr/ports/input"

	"github.com/gofiber/fiber/v2"
)

type AuthenticationAdapter struct {
	authenticationAdapter input.AuthenticationPort
}

func NewAuthenticationAdapter(authenticationAdapter input.AuthenticationPort) *AuthenticationAdapter {
	return &AuthenticationAdapter{authenticationAdapter: authenticationAdapter}
}

func (a *AuthenticationAdapter) Login(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	if err := a.authenticationAdapter.Login(&user); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Logged In Successfully",
	})
}


func (a *AuthenticationAdapter) Logout(c *fiber.Ctx) error {
	if err := a.authenticationAdapter.Logout(); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Error while logging out",
        })
    }
    return c.JSON(fiber.Map{
        "message": "Logged Out Successfully",
    })
}

func (a *AuthenticationAdapter) IsEmailExist(c *fiber.Ctx) error {
	email := c.Query("email")
    return c.JSON(fiber.Map{
        "isEmailExist": a.authenticationAdapter.IsEmailExist(email),
    })
}
