package outputPort

import "backend-food-menu-qr/core/domain"

type OrderOutputPort interface {
	UserOrderOutputPort
	OwnerOrderOutputPort
	GetOrderByOrderId(restaurantId string, orderId string) (*domain.Order, error)
	SaveOrder(restaurantId string, order *domain.Order) (*domain.Order, error)
}

type UserOrderOutputPort interface {
	DeleteOrder(restaurantId string, orderID string) error
	GetOrderByUserIdDateAndStatus(userID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error)
}

type OwnerOrderOutputPort interface {
	GetOrderByRestaurantIdDateAndStatus(restaurantID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error)
}
