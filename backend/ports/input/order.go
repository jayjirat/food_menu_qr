package inputPort

import "backend-food-menu-qr/core/domain"

type OrderInputPort interface {
	UserOrderInputPort
	OwnerOrderInputPort
	GetOrderByID(orderId string) (*domain.Order, error)
}

type UserOrderInputPort interface {
	CreateOrder(order *domain.Order) (*domain.Order, error) //-//
	UpdateOrder(orderID string, order *domain.Order) (*domain.Order, error)
	DeleteOrder(orderID string) error

	GetOrderByUserID(userID string) ([]*domain.Order, error)
	GetOrderByUserIdDateAndStatus(userID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error)
}

type OwnerOrderInputPort interface {
	UpdateOrderStatus(orderID string, status domain.OrderStatus) (*domain.Order, error) //-//

	GetOrdersByRestaurantID(restaurantID string) ([]*domain.Order, error)                                                                   //-//
	GetOrderByRestaurantIDDateAndStatus(restaurantID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error) //-//
}
