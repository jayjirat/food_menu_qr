package inputPort

import "backend-food-menu-qr/core/domain"

type FoodInputPort interface {
	CreateFood(food *domain.Food) (*domain.Food, error)
	UpdateFood(food *domain.Food) (*domain.Food, error)
	DeleteFood(food *domain.Food) error
	GetFoodByID(id string) (*domain.Food, error)
	GetAllFoodsByRestaurantID(restaurantId string) ([]*domain.Food, error)
}
