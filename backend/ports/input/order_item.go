package inputPort

import "backend-food-menu-qr/core/domain"

type OrderItemInputPort interface {
	UserOrderItemInputPort
}

type UserOrderItemInputPort interface {
	AddItemToOrder(item *domain.OrderItem) (*domain.OrderItem, error)
	RemoveItemFromOrder(itemID string) error
}
