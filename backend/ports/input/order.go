package inputPort

import "backend-food-menu-qr/core/domain"

type OrderInputPort interface {
	UserOrderInputPort
	OwnerOrderInputPort
	GetOrderByOrderId(restaurantId string,orderId string) (*domain.Order, error)
}

type UserOrderInputPort interface {
	CreateOrder(restaurantId string,order *domain.Order) (*domain.Order, error)
	UpdateOrder(restaurantId string,orderID string, order *domain.Order) (*domain.Order, error)
	DeleteOrder(restaurantId string,orderID string) error
	GetOrderByUserIdDateAndStatus(userID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error)
}

type OwnerOrderInputPort interface {
	UpdateOrderStatus(restaurantId string,orderID string, status domain.OrderStatus) (*domain.Order, error)
	GetOrderByRestaurantIdDateAndStatus(restaurantID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error)
}
