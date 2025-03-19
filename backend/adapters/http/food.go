package httpAdapter

import (
	"backend-food-menu-qr/core/domain"
	inputPort "backend-food-menu-qr/ports/input"

	"github.com/gofiber/fiber/v2"
)

type FoodInputAdapter struct {
	foodInputPort inputPort.FoodInputPort
}

func NewFoodInputAdapter(foodInputPort inputPort.FoodInputPort) *FoodInputAdapter {
	return &FoodInputAdapter{foodInputPort: foodInputPort}
}

func (f *FoodInputAdapter) CreateFood(c *fiber.Ctx) error {
	var food domain.Food
	var restaurantId = c.Params("restaurantId")
	if restaurantId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID is required",
		})
	}
	if err := c.BodyParser(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if food.RestaurantID == "" || food.Name == "" || food.Price == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid food data",
		})
	}
	newFood, err := f.foodInputPort.CreateFood(restaurantId, &food)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newFood)
}

func (f *FoodInputAdapter) UpdateFood(c *fiber.Ctx) error {
	var food domain.Food
	var restaurantId = c.Params("restaurantId")
	if restaurantId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID is required",
		})
	}

	if err := c.BodyParser(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	updatedFood, err := f.foodInputPort.UpdateFood(restaurantId, &food)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedFood)
}

func (f *FoodInputAdapter) DeleteFood(c *fiber.Ctx) error {
	var restaurantId = c.Params("restaurantId")
	var foodId = c.Params("foodId")

	if restaurantId == "" || foodId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID and food ID are required",
		})
	}

	err := f.foodInputPort.DeleteFood(restaurantId, foodId)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Food deleted successfully",
	})
}

func (f *FoodInputAdapter) GetFoodByRestaurantIdAndFoodId(c *fiber.Ctx) error {
	var restaurantId = c.Params("restaurantId")
	var foodId = c.Params("foodId")
	if restaurantId == "" || foodId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Restaurant ID and food ID are required",
		})
	}
	foods, err := f.foodInputPort.GetFoodByRestaurantIdAndFoodId(restaurantId, foodId)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON((fiber.Map{
			"message": err.Error(),
		}))
	}

	return c.Status(fiber.StatusOK).JSON(foods)
}

func (f *FoodInputAdapter) GetAllFoodsByRestaurantID(c *fiber.Ctx) error {
	var restaurantId = c.Params("restaurantId")
	foods, err := f.foodInputPort.GetAllFoodsByRestaurantId(restaurantId)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(foods)
}
