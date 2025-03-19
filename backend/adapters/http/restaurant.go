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

	if restaurant.OwnerID == "" || restaurant.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid restaurant data",
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
	restaurantId := c.Params("restaurantId")
	if restaurantId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID is required",
		})
	}
	if err := c.BodyParser(&restaurant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}

	if restaurant.ID != restaurantId {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID mismatch",
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
	restaurantId := c.Params("restaurantId")
	if restaurantId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID is required",
		})
	}
	if err := r.restaurantInputPort.DeleteRestaurant(restaurantId); err != nil {
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
	userId := c.Query("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User ID is required",
		})
	}
	restaurant, err := r.restaurantInputPort.GetMyRestaurant(userId)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(restaurant)
}

func (r *RestaurantInputAdapter) GetRestaurantByID(c *fiber.Ctx) error {
	restaurantId := c.Params("restaurantId")

	// handle query parameters error
	if restaurantId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID is required",
		})
	}

	restaurant, err := r.restaurantInputPort.GetRestaurantByID(restaurantId)
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

func (r *RestaurantInputAdapter) OwnerUpdateRestaurantStatus(c *fiber.Ctx) error {
	var rs domain.RestaurantStatus
	restaurantId := c.Params("restaurantId")
	if restaurantId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID is required",
		})
	}
	var updateRestaurantStatusRequest domain.UpdateRestaurantStatusRequest

	if err := c.BodyParser(&updateRestaurantStatusRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}
	restaurant, err := r.restaurantInputPort.OwnerUpdateRestaurantStatus(restaurantId, rs.ToRestaurantStatus(updateRestaurantStatusRequest.Status))
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(restaurant)
}

func (r *RestaurantInputAdapter) AdminUpdateRestaurantStatus(c *fiber.Ctx) error {
	var rs domain.RestaurantStatus
	restaurantId := c.Params("restaurantId")
	if restaurantId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID is required",
		})
	}
	var updateRestaurantStatusRequest domain.UpdateRestaurantStatusRequest

	if err := c.BodyParser(&updateRestaurantStatusRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}

	adminUpdatedStatus := rs.ToRestaurantStatus(updateRestaurantStatusRequest.Status)

	if adminUpdatedStatus != domain.RestaurantStatusInactive {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Only admin can update restaurant status to Inactive",
		})
	}
	restaurant, err := r.restaurantInputPort.AdminUpdateRestaurantStatus(restaurantId, adminUpdatedStatus)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(restaurant)
}
