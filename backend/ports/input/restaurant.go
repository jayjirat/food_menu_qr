package inputPort

import "backend-food-menu-qr/core/domain"

type UserRestaurantInputPort interface {
	GetRestaurantByID(restaurantId string) (*domain.Restaurant, error)
}

type OwnerRestaurantInputPort interface {
	CreateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	UpdateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	DeleteRestaurant(restaurantId string) error
	GetMyRestaurant(userId string) ([]*domain.Restaurant, error)
	OwnerUpdateRestaurantStatus(restaurantId string, status domain.RestaurantStatus) (*domain.Restaurant, error)
}

type AdminRestaurantInputPort interface {
	GetAllRestaurants() ([]*domain.Restaurant, error)
	AdminUpdateRestaurantStatus(restaurantId string, status domain.RestaurantStatus) (*domain.Restaurant, error)
}

type RestaurantInputPort interface {
	UserRestaurantInputPort
	OwnerRestaurantInputPort
	AdminRestaurantInputPort
}
