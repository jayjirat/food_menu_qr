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

func (o *OrderOutputAdapter) GetOrderByOrderId(restaurantId string, orderId string) (*domain.Order, error) {
	var order domain.Order
	if err := o.db.Where("restaurant_id =? AND id =?", restaurantId, orderId).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderOutputAdapter) SaveOrder(restaurantId string, order *domain.Order) (*domain.Order, error) {
	var existingOrder domain.Order
	if err := o.db.Where("restaurant_id =? AND id =?", restaurantId, order.ID).First(&existingOrder).Error; err == nil {
		if err := o.db.Model(&existingOrder).Updates(order).Error; err != nil {
			return nil, err
		}
		return &existingOrder, nil
	}
	err := o.db.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderOutputAdapter) DeleteOrder(restaurantId string, orderID string) error {
	if err := o.db.Where("restaurant_id =? AND id =?", restaurantId, orderID).Delete(&domain.Order{}).Error; err != nil {
		return err
	}
	return nil
}

func (o *OrderOutputAdapter) GetOrderByUserIdDateAndStatus(userID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error) {
	var orders []*domain.Order
	query := o.db.Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	} else if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	} else if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderOutputAdapter) GetOrderByRestaurantIdDateAndStatus(restaurantID string, startDate, endDate string, status domain.OrderStatus) ([]*domain.Order, error) {
	var orders []*domain.Order
	query := o.db.Where("restaurant_id =?", restaurantID)

	if status != "" {
		query = query.Where("status =?", status)
	}

	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN? AND?", startDate, endDate)
	} else if startDate != "" {
		query = query.Where("created_at >=?", startDate)
	} else if endDate != "" {
		query = query.Where("created_at <=?", endDate)
	}

	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
