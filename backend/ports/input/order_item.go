package inputPort

import "backend-food-menu-qr/core/domain"

type OrderItemInputPort interface {
	AddItemToOrder(orderID string, item *domain.OrderItem) (*domain.OrderItem, error)
	RemoveItemFromOrder(orderID string, itemID string) error
	DeleteOrder(order *domain.Order) error
}
