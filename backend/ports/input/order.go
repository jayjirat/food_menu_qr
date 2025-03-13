package inputPort

import "backend-food-menu-qr/core/domain"

type OrderInputPort interface {

	// Owner
	GetOrderByID(id string) (*domain.Order, error)
	GetOrdersByRestaurant(restaurantID string) ([]*domain.Order, error)
	UpdateOrderStatus(orderID string, status domain.OrderStatus) error
	
	GetOrderHistoryByRestaurantID(restaurantID string) ([]*domain.Order, error)
	GetOrderHistoryByRestaurantIDAndStatus(restaurantID string,status domain.OrderStatus) ([]*domain.Order, error)

	// user
	CreateOrder(order *domain.Order) (*domain.Order, error)
	UpdateOrder(order *domain.Order) (*domain.Order, error)
	GetOrderHistoryByUserID(userID string) ([]*domain.Order, error)
	GetOrderHistoryByUserIDAndStatus(userID string,status domain.OrderStatus) ([]*domain.Order, error)
	
	// all
	GetOrderHistoryByDate(restaurantID string, startDate, endDate string) ([]*domain.Order, error)
	
	
	// service -> CalculateTotalPrice(order *domain.Order) (float64, error)
}
