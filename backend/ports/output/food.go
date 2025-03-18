package outputPort

import "backend-food-menu-qr/core/domain"

type FoodOutputPort interface {
	OwnerFoodOutputPort
}

type OwnerFoodOutputPort interface {
	SaveFood(food *domain.Food) (*domain.Food, error)
	DeleteFood(food *domain.Food) error
	GetFoodByID(id string) (*domain.Food, error)
	GetAllFoodsByRestaurantID(restaurantId string) ([]*domain.Food, error)
}
