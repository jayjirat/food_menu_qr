// First file: food_and_restaurant.go
package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestaurantStatus string

const (
	RestaurantStatusOpene    RestaurantStatus = "Open"
	RestaurantStatusInactive RestaurantStatus = "Inactive"
	RestaurantStatusClose    RestaurantStatus = "Close"
)

type UpdateRestaurantStatusRequest struct {
	Status string `json:"status"`
}

type FoodCategory string

const (
	CategoryMainCourse FoodCategory = "Main Course"
	CategoryDessert    FoodCategory = "Dessert"
	CategorySnack      FoodCategory = "Snack"
	CategoryDrink      FoodCategory = "Drink"
	CategoryBakery     FoodCategory = "Bakery"
	CategoryOther      FoodCategory = "Other"
)

type UpdateFoodCategoryRequest struct {
	Status string `json:"status"`
}

type OrderStatus string

const (
	OrderStatusActive    OrderStatus = "Active"
	OrderStatusConfirmed OrderStatus = "Confirmed"
	OrderStatusCanceled  OrderStatus = "Canceled"
)

type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}

func (os *OrderStatus) ToOrderStatus(orderStatus string) OrderStatus {
	switch orderStatus {
	case string(OrderStatusActive):
		return OrderStatusConfirmed
	case string(OrderStatusConfirmed):
		return OrderStatusCanceled
	case string(OrderStatusCanceled):
		return OrderStatusActive
	default:
		return OrderStatusActive
	}
}

type TableStatus string

const (
	TableStatusAvailable TableStatus = "Available"
	TableStatusOccupied  TableStatus = "Occupied"
	TableStatusReserved  TableStatus = "Reserved"
)

type Restaurant struct {
	ID      string           `gorm:"type:uuid;primaryKey" json:"id"`
	Name    string           `json:"name"`
	LogoUrl string           `json:"logoUrl"`
	Foods   []Food           `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"foods"`
	Orders  []Order          `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"orders"`
	Tables  []Table          `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"tables"`
	Status  RestaurantStatus `json:"status"`
}

func (r *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.NewString()
	return nil
}

type Food struct {
	ID           string       `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID string       `gorm:"type:uuid;not null;index" json:"restaurantId"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Price        float64      `json:"price"`
	ImageUrl     string       `json:"imageUrl"`
	Category     FoodCategory `json:"category"`
}

func (f *Food) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.NewString()
	return nil
}

type Order struct {
	ID           string      `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID string      `gorm:"type:uuid;not null" json:"restaurantId"`
	TableID      string      `gorm:"type:uuid;index" json:"tableId"`
	Table        Table       `gorm:"foreignKey:TableID;references:ID" json:"table"`
	UserID       string      `gorm:"type:uuid;not null;index" json:"userId"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"orderItems"`
	TotalPrice   float64     `json:"totalPrice"`
	Status       OrderStatus `json:"status"`
	OrderTime    time.Time   `json:"orderTime"`
	TakeAway     bool        `json:"takeAway"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.NewString()
	return nil
}

type OrderItem struct {
	ID            string  `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID       string  `gorm:"type:uuid;not null" json:"orderId"`
	FoodID        string  `gorm:"type:uuid;index" json:"foodId"`
	Food          Food    `gorm:"foreignKey:FoodID;references:ID" json:"food"`
	Quantity      int     `json:"quantity"`
	SubTotalPrice float64 `json:"subTotalPrice"`
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	oi.ID = uuid.NewString()
	return nil
}

type Table struct {
	ID           string      `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID string      `gorm:"type:uuid;not null;index" json:"restaurantId"`
	Number       int         `json:"number"`
	Status       TableStatus `json:"status"`
}

func (t *Table) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return nil
}
