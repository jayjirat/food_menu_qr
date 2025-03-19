package usecase

import (
	"backend-food-menu-qr/core/domain"
	outputPort "backend-food-menu-qr/ports/output"
)

type FoodUseCase struct {
	foodOutputPort       outputPort.FoodOutputPort
	restaurantOutputPort outputPort.RestaurantOutputPort
}

func NewFoodUseCase(foodOutputPort outputPort.FoodOutputPort, restaurantOutputPort outputPort.RestaurantOutputPort) *FoodUseCase {
	return &FoodUseCase{foodOutputPort: foodOutputPort, restaurantOutputPort: restaurantOutputPort}
}

func (f *FoodUseCase) CreateFood(restaurantId string, food *domain.Food) (*domain.Food, error) {
	if _, err := f.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return nil, err
	}
	food.CreatedAt = domain.GetCurrentTime()
	food.UpdatedAt = domain.GetCurrentTime()
	return f.foodOutputPort.SaveFood(food)
}

func (f *FoodUseCase) UpdateFood(restaurantId string, food *domain.Food) (*domain.Food, error) {
	if _, err := f.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return nil, err
	}

	updatedFood, err := f.foodOutputPort.GetFoodByRestaurantIdAndFoodId(restaurantId, food.ID)
	if err != nil {
		return nil, err
	}

	if food.Name != "" {
		updatedFood.Name = food.Name
	}
	if food.Description != "" {
		updatedFood.Description = food.Description
	}
	if food.Price != 0 {
		updatedFood.Price = food.Price
	}
	if food.ImageUrl != "" {
		updatedFood.ImageUrl = food.ImageUrl
	}
	updatedFood.Category = food.Category
	updatedFood.UpdatedAt = domain.GetCurrentTime()
	return f.foodOutputPort.SaveFood(updatedFood)
}

func (f *FoodUseCase) DeleteFood(restaurantId string, foodId string) error {
	if _, err := f.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return err
	}

	food, err := f.foodOutputPort.GetFoodByRestaurantIdAndFoodId(restaurantId, foodId)
	if err != nil {
		return err
	}

	return f.foodOutputPort.DeleteFood(food)
}

func (f *FoodUseCase) GetFoodByRestaurantIdAndFoodId(restaurantId string, foodId string) (*domain.Food, error) {
	if _, err := f.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return nil, err
	}

	return f.foodOutputPort.GetFoodByRestaurantIdAndFoodId(restaurantId, foodId)
}

func (f *FoodUseCase) GetAllFoodsByRestaurantId(restaurantId string) ([]*domain.Food, error) {
	if _, err := f.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return nil, err
	}

	return f.foodOutputPort.GetAllFoodsByRestaurantId(restaurantId)
}
