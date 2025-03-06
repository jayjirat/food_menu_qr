package http

import (
	"backend-food-menu-qr/core/domain"
	"backend-food-menu-qr/ports/input"

	"github.com/gofiber/fiber/v2"
)

type UserInputAdapter struct {
	userInputPort input.UserInputPort
}

func NewUserInputAdapter(userInputPort input.UserInputPort) *UserInputAdapter {
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

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing JSON"})
	}

	updatedUser, err := u.userInputPort.UpdateUser(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func (u *UserInputAdapter) GetUserByID(c *fiber.Ctx) error {

	id := c.Params("id")
	user, err := u.userInputPort.GetUserByID(id)

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
