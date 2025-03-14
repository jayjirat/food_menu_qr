package inputPort

import "backend-food-menu-qr/core/domain"

type OrderInputPort interface {

	// Owner
	GetOrderByID(orderId string) (*domain.Order, error)
	GetOrdersByRestaurantID(restaurantID string) ([]*domain.Order, error)
	UpdateOrderStatus(orderID string, status domain.OrderStatus) (*domain.Order, error)

	GetOrderByRestaurantIDAndStatus(restaurantID string, status domain.OrderStatus) ([]*domain.Order, error)

	// user
	CreateOrder(order *domain.Order) (*domain.Order, error)
	UpdateOrder(orderID string, order *domain.Order) (*domain.Order, error)
	DeleteOrder(orderID string) error
	
	GetOrderByUserID(userID string) ([]*domain.Order, error)
	GetOrderByUserIDAndStatus(userID string, status domain.OrderStatus) ([]*domain.Order, error)

	// all
	GetOrderByDate(restaurantID string, startDate, endDate string) ([]*domain.Order, error)

	// service -> CalculateTotalPrice(order *domain.Order) (float64, error)
}
