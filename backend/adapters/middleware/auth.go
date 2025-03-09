package middleware

import (
	"backend-food-menu-qr/config"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateToken(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Access Denied: No token provided",
		})
	}

	token = strings.Replace(token, "Bearer ", "", 1)

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token method")
		}
		return []byte(config.AppConfig.SECRET_KEY_JWT), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Access Denied: Invalid token",
		})
	}

	return c.Next()
}
