package inputPort

import "backend-food-menu-qr/core/domain"

type RestaurantInputPort interface {
	// Owner
	CreateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	UpdateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	DeleteRestaurant(restaurantId string) error
	GetMyRestaurant(userId string) ([]*domain.Restaurant, error)
	GetRestaurantByID(restaurantId string) (*domain.Restaurant, error)

	// Admin
	GetAllRestaurants() ([]*domain.Restaurant, error)
}
