package usecase

import (
	"backend-food-menu-qr/core/domain"
	outputPort "backend-food-menu-qr/ports/output"
	"errors"
)

type RestaurantUseCase struct {
	restaurantOutputPort outputPort.RestaurantOutputPort
	userOutputPort       outputPort.UserOutputPort
}

func NewRestaurantUseCase(restaurantOutputPort outputPort.RestaurantOutputPort, userOutputPort outputPort.UserOutputPort) *RestaurantUseCase {
	return &RestaurantUseCase{restaurantOutputPort: restaurantOutputPort, userOutputPort: userOutputPort}
}

func (r *RestaurantUseCase) CreateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error) {
	if _, err := r.userOutputPort.GetUserByUserId(restaurant.OwnerID); err != nil {
		return nil, errors.New("user not found")
	}
	restaurant.CreatedAt = domain.GetCurrentTime()
	restaurant.UpdatedAt = domain.GetCurrentTime()
	newRestaurant, err := r.restaurantOutputPort.SaveRestaurant(restaurant)
	if err != nil {
		return nil, err
	}

	return newRestaurant, nil
}

func (r *RestaurantUseCase) UpdateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error) {
	updatedRestaurant, err := r.restaurantOutputPort.GetRestaurantByID(restaurant.ID)
	if err != nil {
		return nil, err
	}

	if updatedRestaurant == nil {
		return nil, errors.New("restaurant not found")
	}

	if restaurant.Name != "" {
		updatedRestaurant.Name = restaurant.Name
	}
	if restaurant.LogoUrl != "" {
		updatedRestaurant.LogoUrl = restaurant.LogoUrl
	}

	updatedRestaurant.UpdatedAt = domain.GetCurrentTime()

	updatedRestaurant, err = r.restaurantOutputPort.SaveRestaurant(updatedRestaurant)

	if err != nil {
		return nil, err
	}
	return updatedRestaurant, nil
}

func (r *RestaurantUseCase) DeleteRestaurant(restaurantId string) error {
	updatedRestaurant, err := r.restaurantOutputPort.GetRestaurantByID(restaurantId)
	if err != nil {
		return err
	}

	if updatedRestaurant == nil {
		return errors.New("restaurant not found")
	}
	return r.restaurantOutputPort.DeleteRestaurant(restaurantId)
}

func (r *RestaurantUseCase) GetMyRestaurant(userId string) ([]*domain.Restaurant, error) {

	if _, err := r.userOutputPort.GetUserByUserId(userId); err != nil {
		return nil, errors.New("user not found")
	}

	restaurants, err := r.restaurantOutputPort.GetMyRestaurant(userId)
	if err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r *RestaurantUseCase) GetRestaurantByID(restaurantId string) (*domain.Restaurant, error) {
	restaurant, err := r.restaurantOutputPort.GetRestaurantByID(restaurantId)
	if err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (r *RestaurantUseCase) GetAllRestaurants() ([]*domain.Restaurant, error) {
	restaurants, err := r.restaurantOutputPort.GetAllRestaurants()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *RestaurantUseCase) OwnerUpdateRestaurantStatus(restaurantId string, status domain.RestaurantStatus) (*domain.Restaurant, error) {
	updatedRestaurant, err := r.restaurantOutputPort.GetRestaurantByID(restaurantId)
	if err != nil {
		return nil, err
	}

	if updatedRestaurant == nil {
		return nil, errors.New("restaurant not found")
	}

	updatedRestaurant.Status = status
	updatedRestaurant.UpdatedAt = domain.GetCurrentTime()
	updatedRestaurant, err = r.restaurantOutputPort.SaveRestaurant(updatedRestaurant)

	if err != nil {
		return nil, err
	}
	return updatedRestaurant, nil
}

func (r *RestaurantUseCase) AdminUpdateRestaurantStatus(restaurantId string, status domain.RestaurantStatus) (*domain.Restaurant, error) {
	updatedRestaurant, err := r.restaurantOutputPort.GetRestaurantByID(restaurantId)
	if err != nil {
		return nil, err
	}

	if updatedRestaurant == nil {
		return nil, errors.New("restaurant not found")
	}

	updatedRestaurant.Status = status
	updatedRestaurant.UpdatedAt = domain.GetCurrentTime()
	updatedRestaurant, err = r.restaurantOutputPort.SaveRestaurant(updatedRestaurant)

	if err != nil {
		return nil, err
	}
	return updatedRestaurant, nil
}
