package outputPort

import "backend-food-menu-qr/core/domain"

type FoodOutputPort interface {
	OwnerFoodOutputPort
}

type OwnerFoodOutputPort interface {
	SaveFood(food *domain.Food) (*domain.Food, error)
	DeleteFood(food *domain.Food) error
	GetFoodByRestaurantIdAndFoodId(restaurantId string, foodId string) (*domain.Food, error)
	GetAllFoodsByRestaurantId(restaurantId string) ([]*domain.Food, error)
}
