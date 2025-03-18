package inputPort

import "backend-food-menu-qr/core/domain"

type OwnerFoodInputPort interface {
	CreateFood(restaurantId string, food *domain.Food) (*domain.Food, error)
	UpdateFood(restaurantId string, food *domain.Food) (*domain.Food, error)
	DeleteFood(restaurantId string, foodId string) error
	GetFoodByRestaurantIdAndFoodId(restaurantId string, foodId string) (*domain.Food, error)
	GetAllFoodsByRestaurantId(restaurantId string) ([]*domain.Food, error)
}

type FoodInputPort interface {
	OwnerFoodInputPort
}
