// First file: food_and_restaurant.go
package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

type RestaurantStatus string

const (
	RestaurantStatusOpen     RestaurantStatus = "Open"
	RestaurantStatusInactive RestaurantStatus = "Inactive"
	RestaurantStatusClose    RestaurantStatus = "Close"
)

type UpdateRestaurantStatusRequest struct {
	Status string `json:"status"`
}

func (rs *RestaurantStatus) ToRestaurantStatus(restaurantStatus string) RestaurantStatus {
	switch restaurantStatus {
	case string(RestaurantStatusOpen):
		return RestaurantStatusOpen
	case string(RestaurantStatusInactive):
		return RestaurantStatusInactive
	case string(RestaurantStatusClose):
		return RestaurantStatusClose
	default:
		return RestaurantStatusOpen
	}
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
		return OrderStatusActive
	case string(OrderStatusConfirmed):
		return OrderStatusConfirmed
	case string(OrderStatusCanceled):
		return OrderStatusCanceled
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
	OwnerID string           `gorm:"type:uuid;index" json:"ownerId"`
	Owner   User             `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" json:"owner"`
	Name    string           `json:"name"`
	LogoUrl string           `json:"logoUrl"`
	Foods   []Food           `json:"foods"`
	Orders  []Order          `json:"orders"`
	Tables  []Table          `json:"tables"`
	Status  RestaurantStatus `json:"status"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

func (r *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.NewString()
	return nil
}

type Food struct {
	ID           string       `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID string       `gorm:"type:uuid;index" json:"restaurantId"`
	Restaurant   Restaurant   `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" json:"restaurant"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Price        float64      `json:"price"`
	ImageUrl     string       `json:"imageUrl"`
	Category     FoodCategory `json:"category"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

func (f *Food) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.NewString()
	return nil
}

type Order struct {
	ID           string      `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID string      `gorm:"type:uuid;index" json:"restaurantId"`
	Restaurant   Restaurant  `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" json:"restaurant"`
	TableID      string      `gorm:"type:uuid;index" json:"tableId"`
	Table        Table       `gorm:"foreignKey:TableID;references:ID" json:"table"`
	UserID       string      `gorm:"type:uuid;index" json:"userId"`
	User         User        `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" json:"user"`
	OrderItems   []OrderItem `json:"orderItems"`
	TotalPrice   float64     `json:"totalPrice"`
	Status       OrderStatus `json:"status"`
	TakeAway     bool        `json:"takeAway"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.NewString()
	return nil
}

type OrderItem struct {
	ID            string  `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID       string  `gorm:"type:uuid;index" json:"orderId"`
	Order         Order   `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" json:"order"`
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
	RestaurantID string      `gorm:"type:uuid;index" json:"restaurantId"`
	Restaurant   Restaurant  `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" json:"restaurant"`
	Number       int         `gorm:"uniqueIndex:idx_restaurant_table" json:"number"`
	Status       TableStatus `json:"status"`
}

func (t *Table) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return nil
}
