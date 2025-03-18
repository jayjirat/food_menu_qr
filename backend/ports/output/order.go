package outputPort

import "backend-food-menu-qr/core/domain"

type OrderOutputPort interface {
	UserOrderOutputPort
	OwnerOrderOutputPort
	SaveOrder(order *domain.Order) (*domain.Order, error)
	GetOrderByDate(restaurantID string, startDate, endDate string) ([]*domain.Order, error)
}

type UserOrderOutputPort interface {
	GetOrderByUserIDAndStatus(userID string, status domain.OrderStatus) ([]*domain.Order, error)
}

type OwnerOrderOutputPort interface {
	GetOrderByID(id string) (*domain.Order, error)
	GetOrderByRestaurantIDAndStatus(restaurantID string, status domain.OrderStatus) ([]*domain.Order, error)
}
