package outputPort

import "backend-food-menu-qr/core/domain"

type OrderItemInputPort interface {
	SaveOrderItem(orderID string, item *domain.OrderItem) (*domain.OrderItem, error)
	DeleteOrder(order *domain.Order) error
}
