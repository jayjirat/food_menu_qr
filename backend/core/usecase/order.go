package usecase

import (
	"backend-food-menu-qr/core/domain"
	outputPort "backend-food-menu-qr/ports/output"
)

type OrderUseCase struct {
	orderOutputPort      outputPort.OrderOutputPort
	restaurantOutputPort outputPort.RestaurantOutputPort
	userOutputPort       outputPort.UserOutputPort
}

func NewOrderUseCase(orderOutputPort outputPort.OrderOutputPort, restaurantOutputPort outputPort.RestaurantOutputPort, userOutputPort outputPort.UserOutputPort) *OrderUseCase {
	return &OrderUseCase{orderOutputPort: orderOutputPort, restaurantOutputPort: restaurantOutputPort, userOutputPort: userOutputPort}
}

func (o *OrderUseCase) GetOrderByOrderId(restaurantId string, orderId string) (*domain.Order, error) {
	if _, err := o.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return nil, err
	}
	return o.orderOutputPort.GetOrderByOrderId(restaurantId, orderId)
}

func (o *OrderUseCase) CreateOrder(restaurantId string, order *domain.Order) (*domain.Order, error) {
	if _, err := o.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return nil, err
	}

	if _, err := o.userOutputPort.GetUserByUserId(order.UserID); err != nil {
		return nil, err
	}
	order.CreatedAt = domain.GetCurrentTime()
	order.UpdatedAt = domain.GetCurrentTime()
	return o.orderOutputPort.SaveOrder(restaurantId, order)
}

func (o *OrderUseCase) UpdateOrder(restaurantId string, orderID string, order *domain.Order) (*domain.Order, error) {
	if _, err := o.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return nil, err
	}

	if _, err := o.userOutputPort.GetUserByUserId(order.UserID); err != nil {
		return nil, err
	}

	updatedOrder, err := o.orderOutputPort.GetOrderByOrderId(restaurantId, orderID)
	if err != nil {
		return nil, err
	}

	if len(order.OrderItems) != len(updatedOrder.OrderItems) {
		updatedOrder.OrderItems = order.OrderItems
	}

	if order.TotalPrice != updatedOrder.TotalPrice {
		updatedOrder.TotalPrice = order.TotalPrice
	}

	if order.TakeAway != updatedOrder.TakeAway {
		updatedOrder.TakeAway = order.TakeAway
	}

	updatedOrder.UpdatedAt = domain.GetCurrentTime()

	return o.orderOutputPort.SaveOrder(restaurantId, updatedOrder)
}

func (o *OrderUseCase) DeleteOrder(restaurantId string, orderID string) error {
	if _, err := o.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return err
	}

	return o.orderOutputPort.DeleteOrder(restaurantId, orderID)
}

func (o *OrderUseCase) GetOrderByUserIdDateAndStatus(userID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error) {
	if _, err := o.userOutputPort.GetUserByUserId(userID); err != nil {
		return nil, err
	}

	return o.orderOutputPort.GetOrderByUserIdDateAndStatus(userID, startDate, endDate, status)
}

func (o *OrderUseCase) UpdateOrderStatus(restaurantId string, orderID string, status domain.OrderStatus) (*domain.Order, error) {
	if _, err := o.restaurantOutputPort.GetRestaurantByID(restaurantId); err != nil {
		return nil, err
	}

	order, err := o.orderOutputPort.GetOrderByOrderId(restaurantId, orderID)
	if err != nil {
		return nil, err
	}

	order.Status = status
	order.UpdatedAt = domain.GetCurrentTime()

	return o.orderOutputPort.SaveOrder(restaurantId, order)
}

func (o *OrderUseCase) GetOrderByRestaurantIdDateAndStatus(restaurantID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error) {
	if _, err := o.restaurantOutputPort.GetRestaurantByID(restaurantID); err != nil {
		return nil, err
	}

	return o.orderOutputPort.GetOrderByRestaurantIdDateAndStatus(restaurantID, startDate, endDate, status)
}
