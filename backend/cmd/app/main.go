package main

import (
	httpAdapter "backend-food-menu-qr/adapters/http"
	"backend-food-menu-qr/adapters/middleware"
	repositoryAdapter "backend-food-menu-qr/adapters/repository"
	"backend-food-menu-qr/config"
	"backend-food-menu-qr/core/domain"
	"backend-food-menu-qr/core/usecase"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	config.LoadConfig()
	appConfig := config.AppConfig
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", appConfig.DB_HOST, appConfig.DB_USER, appConfig.DB_PASSWORD, appConfig.DB_NAME, appConfig.DB_PORT)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&domain.User{}, &domain.Restaurant{}, &domain.Food{}, &domain.Order{}, &domain.OrderItem{}, &domain.Table{})
}

func main() {
	initDB()
	userRepo := repositoryAdapter.NewUserOutputAdapter(db)
	restaurantRepo := repositoryAdapter.NewRestaurantOutputPort(db)
	foodRepo := repositoryAdapter.NewFoodOutputAdapter(db)

	userUsecase := usecase.NewUserUseCase(userRepo)
	authenticationUsecase := usecase.NewAuthenticationUseCase(userRepo)
	restaurantUsecase := usecase.NewRestaurantUseCase(restaurantRepo)
	foodUsecase := usecase.NewFoodUseCase(foodRepo)

	userInputAdapter := httpAdapter.NewUserInputAdapter(userUsecase)
	authenticationInputAdapter := httpAdapter.NewAuthenticationAdapter(authenticationUsecase)
	restaurantInputAdapter := httpAdapter.NewRestaurantInputAdapter(restaurantUsecase)
	foodInputAdapter := httpAdapter.NewFoodInputAdapter(foodUsecase)

	app := fiber.New()

	app.Post("/api/register", authenticationInputAdapter.Register)
	app.Post("/api/login", authenticationInputAdapter.Login)

	app.Use(middleware.AuthenticateToken)

	app.Get("/api/user/:userId", userInputAdapter.GetUserByUserId)
	app.Put("/api/user/:userId", userInputAdapter.UpdateUser)
	app.Get("/api/admin/users", userInputAdapter.GetAllUsers)
	app.Get("/api/admin/owners", userInputAdapter.GetAllOwners)
	app.Post("/api/admin/user", userInputAdapter.CreateUser)
	app.Delete("/api/user/:userId", userInputAdapter.DeleteUser)

	app.Get("/api/owner/restaurant", restaurantInputAdapter.GetMyRestaurant)
	app.Post("/api/owner/restaurant", restaurantInputAdapter.CreateRestaurant)
	app.Put("/api/owner/restaurant/:restaurantId/details", restaurantInputAdapter.UpdateRestaurant)
	app.Put("/api/owner/restaurant/:restaurantId/status", restaurantInputAdapter.OwnerUpdateRestaurantStatus)
	app.Delete("/api/owner/restaurant/:restaurantId", restaurantInputAdapter.DeleteRestaurant)

	app.Get("/api/restaurant/:restaurantId", restaurantInputAdapter.GetRestaurantByID)

	app.Get("/api/admin/restaurants", restaurantInputAdapter.GetAllRestaurants)
	app.Put("/api/admin/restaurant/:restaurantId/status", restaurantInputAdapter.AdminUpdateRestaurantStatus)

	app.Post("/api/owner/restaurant/:restaurantId/food", foodInputAdapter.CreateFood)
	app.Put("/api/owner/restaurant/:restaurantId/food", foodInputAdapter.UpdateFood)
	app.Delete("/api/owner/restaurant/:restaurantId/food/:foodId", foodInputAdapter.DeleteFood)
	app.Get("/api/restaurant/:restaurantId/food/:foodId", foodInputAdapter.GetFoodByRestaurantIdAndFoodId)
	app.Get("/api/restaurant/:restaurantId/foods", foodInputAdapter.GetAllFoodsByRestaurantID)

	app.Listen(":" + config.AppConfig.APIPort)
}
