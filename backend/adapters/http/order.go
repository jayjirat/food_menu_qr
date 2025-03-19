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

func (o *OrderInputAdapter) GetOrderByOrderId(c *fiber.Ctx) error {
	var restaurantId = c.Params("restaurantId")
	var orderId = c.Params("orderId")
	if restaurantId == "" || orderId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID and Order ID is required",
		})
	}
	order, err := o.orderInputPort.GetOrderByOrderId(restaurantId, orderId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (o *OrderInputAdapter) CreateOrder(c *fiber.Ctx) error {
	var restaurantId = c.Params("restaurantId")
	if restaurantId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID is required",
		})
	}
	var order domain.Order
	err := c.BodyParser(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if order.RestaurantID == "" || order.TableID == "" || order.UserID == "" || len(order.OrderItems) == 0 || order.TotalPrice == 0 || order.Status != domain.OrderStatusActive {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid order data",
		})
	}

	newOrder, err := o.orderInputPort.CreateOrder(restaurantId, &order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newOrder)
}

func (o *OrderInputAdapter) UpdateOrder(c *fiber.Ctx) error {
	var restaurantId = c.Params("restaurantId")
	var orderId = c.Params("orderId")
	if restaurantId == "" || orderId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID and Order ID is required",
		})
	}
	var order domain.Order

	err := c.BodyParser(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	updatedOrder, err := o.orderInputPort.UpdateOrder(restaurantId, orderId, &order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedOrder)
}

func (o *OrderInputAdapter) DeleteOrder(c *fiber.Ctx) error {
	var restaurantId = c.Params("restaurantId")
	var orderId = c.Params("orderId")
	if restaurantId == "" || orderId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID and Order ID is required",
		})
	}

	err := o.orderInputPort.DeleteOrder(restaurantId, orderId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order deleted successfully",
	})
}

func (o *OrderInputAdapter) GetOrderByUserIdDateAndStatus(c *fiber.Ctx) error {
	var userId = c.Query("userId")
	var os domain.OrderStatus
	var orderStatus = c.Query("orderStatus")
	var startDate = c.Query("startDate")
	var endDate = c.Query("endDate")

	orders, err := o.orderInputPort.GetOrderByUserIdDateAndStatus(userId, startDate, endDate, os.ToOrderStatus(orderStatus))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}

func (o *OrderInputAdapter) UpdateOrderStatus(c *fiber.Ctx) error {
	var restaurantId = c.Params("restaurantId")
	var orderId = c.Params("orderId")
	if restaurantId == "" || orderId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID and Order ID is required",
		})
	}
	var os domain.OrderStatus
	var updateOrderStatusRequest domain.UpdateOrderStatusRequest

	if err := c.BodyParser(&updateOrderStatusRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	updatedStatusOrder, err := o.orderInputPort.UpdateOrderStatus(restaurantId, orderId, os.ToOrderStatus(updateOrderStatusRequest.Status))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedStatusOrder)
}

func (o *OrderInputAdapter) GetOrderByRestaurantIdDateAndStatus(c *fiber.Ctx) error {
	var os domain.OrderStatus
	var restaurantId = c.Params("restaurantId")
	var orderStatus = c.Query("orderStatus")
	var startDate = c.Query("startDate")
	var endDate = c.Query("endDate")

	orders, err := o.orderInputPort.GetOrderByRestaurantIdDateAndStatus(restaurantId, startDate, endDate, os.ToOrderStatus(orderStatus))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}
