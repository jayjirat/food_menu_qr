package repositoryAdapter

import (
	"backend-food-menu-qr/core/domain"

	"gorm.io/gorm"
)

type OrderItemOutputAdapter struct {
	db *gorm.DB
}

func NewOrderItemOutputAdapter(db *gorm.DB) *OrderItemOutputAdapter {
	return &OrderItemOutputAdapter{db: db}
}

func (o *OrderItemOutputAdapter) SaveOrderItem(orderID string, item *domain.OrderItem) (*domain.OrderItem, error) {
	if err := o.db.Create(&domain.OrderItem{OrderID: orderID, FoodID: item.FoodID, Quantity: item.Quantity}).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (o *OrderItemOutputAdapter) DeleteOrderItem(orderID, foodID string) error {
	if err := o.db.Where("order_id =? AND food_id =?", orderID, foodID).Delete(&domain.OrderItem{}).Error; err != nil {
		return err
	}

	return nil
}
