package httpAdapter

import (
	"backend-food-menu-qr/core/domain"
	inputPort "backend-food-menu-qr/ports/input"

	"github.com/gofiber/fiber/v2"
)

type UserInputAdapter struct {
	userInputPort inputPort.UserInputPort
}

func NewUserInputAdapter(userInputPort inputPort.UserInputPort) *UserInputAdapter {
	return &UserInputAdapter{userInputPort: userInputPort}
}

func (u *UserInputAdapter) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	createdUser, err := u.userInputPort.CreateUser(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (u *UserInputAdapter) UpdateUser(c *fiber.Ctx) error {
	var user domain.User
	userId := c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User ID is required",
		})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing JSON"})
	}

	if user.ID != userId {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User ID mismatch",
		})
	}

	updatedUser, err := u.userInputPort.UpdateUser(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func (u *UserInputAdapter) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
    if userId == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "User ID is required",
        })
    }

    err := u.userInputPort.DeleteUser(userId)

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}

func (u *UserInputAdapter) GetUserByUserId(c *fiber.Ctx) error {

	userId := c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User ID is required"})
	}

	user, err := u.userInputPort.GetUserByUserId(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserInputAdapter) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.userInputPort.GetAllUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (u *UserInputAdapter) GetAllOwners(c *fiber.Ctx) error {
	users, err := u.userInputPort.GetAllOwners()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
