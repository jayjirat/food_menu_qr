package repositoryAdapter

import (
	"backend-food-menu-qr/core/domain"

	"gorm.io/gorm"
)

type FoodOutputAdapter struct {
	db *gorm.DB
}

func NewFoodOutputAdapter(db *gorm.DB) *FoodOutputAdapter {
	return &FoodOutputAdapter{db: db}
}

func (f *FoodOutputAdapter) SaveFood(food *domain.Food) (*domain.Food, error) {
	if err := f.db.Create(food).Error; err != nil {
		return nil, err
	}
	return food, nil
}

func (f *FoodOutputAdapter) DeleteFood(food *domain.Food) error {
	if err := f.db.Delete(food).Error; err != nil {
		return err
	}
	return nil
}

func (f *FoodOutputAdapter) GetFoodByID(id string) (*domain.Food, error) {
	var food domain.Food
	if err := f.db.Where("id =?", id).First(&food).Error; err != nil {
		return nil, err
	}
	return &food, nil
}

func (r *FoodOutputAdapter) GetAllFoodsByRestaurantID(restaurantId string) ([]*domain.Food, error) {
	var foods []*domain.Food
	if err := r.db.Where("restaurant_id =?", restaurantId).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}
