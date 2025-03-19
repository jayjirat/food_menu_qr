package middleware

import (
	"backend-food-menu-qr/config"
	"fmt"
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

	value, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
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

	claims, ok := value.Claims.(jwt.MapClaims)
	if !ok || !value.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Access Denied: Invalid token",
		})
	}

	userId, _ := claims["user_id"].(string)
	role, _ := claims["role"].(string)

	c.Locals("userId", userId)
	c.Locals("role", role)
	fmt.Print(userId)
	return c.Next()
}

func RequireOwnerRole(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "owner" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access Denied: You must be an owner",
		})
	}
	return c.Next()
}

func RequireAdminRole(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role == nil || role.(string) != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access Denied: You must be an admin",
		})
	}
	return c.Next()
}

func RequireSameUserOrAdmin(c *fiber.Ctx) error {
    userId := c.Locals("userId").(string)
    role := c.Locals("role").(string)
    paramUserId := c.Params("userId")

    if role != "admin" && userId != paramUserId {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "message": "Access Denied: You can only access your own data",
        })
    }

    return c.Next()
}

