package repositoryAdapter

import (
	"backend-food-menu-qr/core/domain"

	"gorm.io/gorm"
)

type OrderOutputAdapter struct {
	db *gorm.DB
}

func NewOrderOutputAdapter(db *gorm.DB) *OrderOutputAdapter {
	return &OrderOutputAdapter{db: db}
}

func (o *OrderOutputAdapter) SaveOrder(order *domain.Order) (*domain.Order, error) {
	err := o.db.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderOutputAdapter) GetOrderByDate(restaurantID string, startDate, endDate string) ([]*domain.Order, error) {
	var orders []*domain.Order
	if err := o.db.Where("restaurant_id =? AND createdAt BETWEEN ? AND?", restaurantID, startDate, endDate).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderOutputAdapter) GetOrderByUserIDAndStatus(userID string, status domain.OrderStatus) ([]*domain.Order, error) {
	var orders []*domain.Order
	if err := o.db.Where("user_id =? AND status =?", userID, status).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderOutputAdapter) GetOrderByID(id string) (*domain.Order, error) {
	var order domain.Order
	if err := o.db.Where("id =?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderOutputAdapter) GetOrderByRestaurantIDAndStatus(restaurantID string, status domain.OrderStatus) ([]*domain.Order, error) {
	var orders []*domain.Order
	if err := o.db.Where("restaurant_id =? AND status =?", restaurantID, status).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
