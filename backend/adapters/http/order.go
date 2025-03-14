package httpAdapter

import (
	"backend-food-menu-qr/core/domain"
	inputPort "backend-food-menu-qr/ports/input"

	"github.com/gofiber/fiber/v2"
)

type OrderInputAdapter struct {
	orderInputPort inputPort.OrderInputPort
}

func NewOrderInputAdapter(orderInputPort inputPort.OrderInputPort) *OrderInputAdapter {
	return &OrderInputAdapter{orderInputPort: orderInputPort}
}

func (o *OrderInputAdapter) GetOrderByID(c *fiber.Ctx) error {
	var orderId = c.Query("orderId")
	order, err := o.orderInputPort.GetOrderByID(orderId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (o *OrderInputAdapter) GetOrdersByRestaurantID(c *fiber.Ctx) error {
	var restaurantId = c.Query("restaurantId")
	orders, err := o.orderInputPort.GetOrdersByRestaurantID(restaurantId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}

func (o *OrderInputAdapter) UpdateOrderStatus(c *fiber.Ctx) error {
	var os domain.OrderStatus
	var orderId = c.Params("orderId")
	var updateOrderStatusRequest domain.UpdateOrderStatusRequest

	if err := c.BodyParser(&updateOrderStatusRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	updatedStatusOrder, err := o.orderInputPort.UpdateOrderStatus(orderId, os.ToOrderStatus(updateOrderStatusRequest.Status))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedStatusOrder)
}

func (o *OrderInputAdapter) GetOrderByRestaurantIDAndStatus(c *fiber.Ctx) error {
	var os domain.OrderStatus
	var restaurantId = c.Query("restaurantId")
	var orderStatus = c.Query("orderStatus")

	orders, err := o.orderInputPort.GetOrderByRestaurantIDAndStatus(restaurantId, os.ToOrderStatus(orderStatus))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}

func (o *OrderInputAdapter) CreateOrder(c *fiber.Ctx) error {
	var order domain.Order
	err := c.BodyParser(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	newOrder, err := o.orderInputPort.CreateOrder(&order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newOrder)
}

func (o *OrderInputAdapter) UpdateOrder(c *fiber.Ctx) error {
	var order domain.Order
	var orderId = c.Params("orderId")

	err := c.BodyParser(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	updatedOrder, err := o.orderInputPort.UpdateOrder(orderId, &order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedOrder)
}

func (o *OrderInputAdapter) DeleteOrder(c *fiber.Ctx) error {
	var orderId = c.Params("orderId")
	err := o.orderInputPort.DeleteOrder(orderId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order deleted successfully",
	})
}

func (o *OrderInputAdapter) GetOrderByUserID(c *fiber.Ctx) error {
	var userId = c.Query("userId")
	orders, err := o.orderInputPort.GetOrderByUserID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}

func (o *OrderInputAdapter) GetOrderByUserIDAndStatus(c *fiber.Ctx) error {
	var userId = c.Query("userId")
	var os domain.OrderStatus
	var orderStatus = c.Query("orderStatus")

	orders, err := o.orderInputPort.GetOrderByUserIDAndStatus(userId, os.ToOrderStatus(orderStatus))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}

func (o *OrderInputAdapter) GetOrderByDate(c *fiber.Ctx) error {
	var startDate, endDate string
	var restaurantId = c.Params("restaurantId")
	startDate = c.Query("startDate")
	endDate = c.Query("endDate")

	orders, err := o.orderInputPort.GetOrderByDate(restaurantId, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}
