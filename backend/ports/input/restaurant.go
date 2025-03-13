package inputPort

import "backend-food-menu-qr/core/domain"

type RestaurantInputPort interface {
	// Owner
	CreateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	UpdateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	DeleteRestaurant(restaurant *domain.Restaurant) error
	GetMyRestaurant(user *domain.User) ([]*domain.Restaurant, error)
	GetRestaurantByID(id string) (*domain.Restaurant, error)

	// Admin
	GetAllRestaurants() ([]*domain.Restaurant, error)
}
