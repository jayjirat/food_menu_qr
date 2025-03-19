package middleware

import (
	"backend-food-menu-qr/config"
	outputPort "backend-food-menu-qr/ports/output"
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

// Middleware factory function
func RequireOwnerOfRestaurant(repo outputPort.RestaurantOutputPort) fiber.Handler {

    return func(c *fiber.Ctx) error {
        userId := c.Locals("userId").(string)
        restaurantId := c.Params("restaurantId")

        restaurant, err := repo.GetRestaurantByID(restaurantId)
        if err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Restaurant not found"})
        }

        if restaurant.OwnerID != userId {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Access Denied: Not the owner"})
        }
        return c.Next()
    }
}
