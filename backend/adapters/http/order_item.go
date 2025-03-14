package httpAdapter

import (
	"backend-food-menu-qr/core/domain"
	inputPort "backend-food-menu-qr/ports/input"

	"github.com/gofiber/fiber/v2"
)

type orderItemInputAdapter struct {
	orderItemInputPort inputPort.OrderItemInputPort
}

func NewOrderItemInputAdapter(orderItemInputPort inputPort.OrderItemInputPort) *orderItemInputAdapter {
	return &orderItemInputAdapter{orderItemInputPort: orderItemInputPort}
}

func (oi *orderItemInputAdapter) AddItemToOrder(c *fiber.Ctx) error {
	var item *domain.OrderItem
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}

	newItem, err := oi.orderItemInputPort.AddItemToOrder(item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newItem)
}

func (oi *orderItemInputAdapter) RemoveItemFromOrder(c *fiber.Ctx) error {
	var itemId = c.Params("itemId")

	err := oi.orderItemInputPort.RemoveItemFromOrder(itemId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item removed successfully"})
}
