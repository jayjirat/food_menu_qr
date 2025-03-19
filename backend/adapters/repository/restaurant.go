package repositoryAdapter

import (
	"backend-food-menu-qr/core/domain"

	"gorm.io/gorm"
)

type RestaurantOutputAdapter struct {
	db *gorm.DB
}

func NewRestaurantOutputPort(db *gorm.DB) *RestaurantOutputAdapter {
	return &RestaurantOutputAdapter{db: db}
}

func (r *RestaurantOutputAdapter) SaveRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error) {
	var existing domain.Restaurant
	if err := r.db.Where("id = ?", restaurant.ID).First(&existing).Error; err == nil {
		if err := r.db.Model(&existing).Updates(restaurant).Error; err != nil {
			return nil, err
		}
		return &existing, nil
	}

	if err := r.db.Create(restaurant).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}


func (r *RestaurantOutputAdapter) DeleteRestaurant(restaurantId string) error {
	var restaurant domain.Restaurant
	if err := r.db.Where("id =?", restaurantId).Delete(&restaurant).Error; err != nil {
		return err
	}

	return nil
}

func (r *RestaurantOutputAdapter) GetMyRestaurant(userId string) ([]*domain.Restaurant, error) {
	var restaurants []*domain.Restaurant
	if err := r.db.Where("ownerId =?", userId).Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r *RestaurantOutputAdapter) GetRestaurantByID(restaurantId string) (*domain.Restaurant, error) {
	var restaurant domain.Restaurant
	if err := r.db.Where("id =?", restaurantId).First(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (r *RestaurantOutputAdapter) GetAllRestaurants() ([]*domain.Restaurant, error) {
	var restaurants []*domain.Restaurant
	if err := r.db.Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}
