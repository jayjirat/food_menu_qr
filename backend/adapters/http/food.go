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
	if err := c.BodyParser(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
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
	var restaurantId = c.Query("restaurantId")
	var foodId = c.Query("foodId")
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
	var restaurantId = c.Query("restaurantId")
	foods, err := f.foodInputPort.GetAllFoodsByRestaurantId(restaurantId)
	if err != nil {
		// TODO: handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(foods)
}
