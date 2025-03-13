package outputPort

import "backend-food-menu-qr/core/domain"

type OrderInputPort interface {
	SaveOrder(order *domain.Order) (*domain.Order, error)

	// Owner
	GetOrderByID(id string) (*domain.Order, error)
	GetOrdersByRestaurant(restaurantID string) ([]*domain.Order, error)

	GetOrderHistoryByRestaurantID(restaurantID string) ([]*domain.Order, error)
	GetOrderHistoryByRestaurantIDAndStatus(restaurantID string, status domain.OrderStatus) ([]*domain.Order, error)

	// user
	GetOrderHistoryByUserID(userID string) ([]*domain.Order, error)
	GetOrderHistoryByUserIDAndStatus(userID string, status domain.OrderStatus) ([]*domain.Order, error)

	// all
	GetOrderHistoryByDate(restaurantID string, startDate, endDate string) ([]*domain.Order, error)
}
