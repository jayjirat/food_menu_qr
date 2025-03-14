package httpAdapter

import (
	"backend-food-menu-qr/core/domain"
	inputPort "backend-food-menu-qr/ports/input"

	"github.com/gofiber/fiber/v2"
)

type RestaurantInputAdapter struct {
	restaurantInputPort inputPort.RestaurantInputPort
}

func NewRestaurantInputAdapter(restaurantInputPort inputPort.RestaurantInputPort) *RestaurantInputAdapter {
	return &RestaurantInputAdapter{restaurantInputPort: restaurantInputPort}
}

func (r *RestaurantInputAdapter) CreateRestaurant(c *fiber.Ctx) error {
	var restaurant domain.Restaurant
	if err := c.BodyParser(&restaurant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	newRestaurant, err := r.restaurantInputPort.CreateRestaurant(&restaurant)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newRestaurant)
}

func (r *RestaurantInputAdapter) UpdateRestaurant(c *fiber.Ctx) error {
	var restaurant domain.Restaurant
	if err := c.BodyParser(&restaurant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	updatedRestaurant, err := r.restaurantInputPort.UpdateRestaurant(&restaurant)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(updatedRestaurant)
}

func (r *RestaurantInputAdapter) DeleteRestaurant(c *fiber.Ctx) error {
	restaurantID := c.Params("restaurantId")
	if err := r.restaurantInputPort.DeleteRestaurant(restaurantID); err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Restaurant deleted successfully",
	})
}

func (r *RestaurantInputAdapter) GetMyRestaurant(c *fiber.Ctx) error {
	userID := c.Query("userId")
	restaurant, err := r.restaurantInputPort.GetMyRestaurant(userID)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(restaurant)
}

func (r *RestaurantInputAdapter) GetRestaurantByID(c *fiber.Ctx) error {
	restaurantID := c.Query("restaurantId")
	restaurant, err := r.restaurantInputPort.GetRestaurantByID(restaurantID)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(restaurant)
}

func (r *RestaurantInputAdapter) GetAllRestaurants(c *fiber.Ctx) error {
	restaurants, err := r.restaurantInputPort.GetAllRestaurants()
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(restaurants)
}
