package outputPort

import "backend-food-menu-qr/core/domain"

type RestaurantOutputPort interface {
	OwnerRestaurantOutputPort
	AdminRestaurantOutputPort
}

type OwnerRestaurantOutputPort interface {
	SaveRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	DeleteRestaurant(restaurantId string) error
	GetMyRestaurant(userId string) ([]*domain.Restaurant, error)
	GetRestaurantByID(restaurantId string) (*domain.Restaurant, error)
}

type AdminRestaurantOutputPort interface {
	GetAllRestaurants() ([]*domain.Restaurant, error)
}
