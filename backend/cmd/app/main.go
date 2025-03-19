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
	orderRepo := repositoryAdapter.NewOrderOutputAdapter(db)

	userUsecase := usecase.NewUserUseCase(userRepo)
	authenticationUsecase := usecase.NewAuthenticationUseCase(userRepo)
	restaurantUsecase := usecase.NewRestaurantUseCase(restaurantRepo, userRepo)
	foodUsecase := usecase.NewFoodUseCase(foodRepo, restaurantRepo)
	orderUsecase := usecase.NewOrderUseCase(orderRepo, restaurantRepo, userRepo)

	userInputAdapter := httpAdapter.NewUserInputAdapter(userUsecase)
	authenticationInputAdapter := httpAdapter.NewAuthenticationAdapter(authenticationUsecase)
	restaurantInputAdapter := httpAdapter.NewRestaurantInputAdapter(restaurantUsecase)
	foodInputAdapter := httpAdapter.NewFoodInputAdapter(foodUsecase)
	orderInputAdapter := httpAdapter.NewOrderInputAdapter(orderUsecase)

	app := fiber.New()

	app.Post("/api/register", authenticationInputAdapter.Register)
	app.Post("/api/login", authenticationInputAdapter.Login)

	app.Use(middleware.AuthenticateToken)

	app.Get("/api/restaurant/:restaurantId", restaurantInputAdapter.GetRestaurantByID)

	app.Get("/api/restaurant/:restaurantId/food/:foodId", foodInputAdapter.GetFoodByRestaurantIdAndFoodId)
	app.Get("/api/restaurant/:restaurantId/foods", foodInputAdapter.GetAllFoodsByRestaurantID)

	app.Get("/api/restaurant/:restaurantId/order/:orderId", orderInputAdapter.GetOrderByOrderId)

	userRoutes := app.Group("/api/user")
	userRoutes.Get("/me", userInputAdapter.GetUserByUserId)
	userRoutes.Put("/me", userInputAdapter.UpdateUser)
	userRoutes.Delete("/me", userInputAdapter.DeleteUser)

	userRoutes.Get("/me/orders", orderInputAdapter.GetOrderByRestaurantIdDateAndStatus)
	userRoutes.Post("/restaurant/:restaurantId/order", orderInputAdapter.CreateOrder)
	userRoutes.Put("/restaurant/:restaurantId/order/:orderId", orderInputAdapter.UpdateOrder)
	userRoutes.Delete("/restaurant/:restaurantId/order/:orderId", orderInputAdapter.DeleteOrder)

	ownerRoutes := app.Group("/api/owner")
	ownerRoutes.Use(middleware.RequireOwnerRole)
	ownerRoutes.Get("/restaurant", restaurantInputAdapter.GetMyRestaurant)
	ownerRoutes.Post("/restaurant", restaurantInputAdapter.CreateRestaurant)

	ownerActionRoutes := app.Group("/api/owner/actions")
	ownerActionRoutes.Use(middleware.RequireOwnerOfRestaurant(restaurantRepo))

	ownerActionRoutes.Put("/restaurant/:restaurantId/details", restaurantInputAdapter.UpdateRestaurant)
	ownerActionRoutes.Patch("/restaurant/:restaurantId/status", restaurantInputAdapter.OwnerUpdateRestaurantStatus)
	ownerActionRoutes.Delete("/restaurant/:restaurantId", restaurantInputAdapter.DeleteRestaurant)

	ownerActionRoutes.Post("/restaurant/:restaurantId/food", foodInputAdapter.CreateFood)
	ownerActionRoutes.Put("/restaurant/:restaurantId/food", foodInputAdapter.UpdateFood)
	ownerActionRoutes.Delete("/restaurant/:restaurantId/food/:foodId", foodInputAdapter.DeleteFood)

	ownerActionRoutes.Patch("/restaurant/:restaurantId/order/:orderId", orderInputAdapter.UpdateOrderStatus)
	ownerActionRoutes.Get("/restaurant/:restaurantId/orders", orderInputAdapter.GetOrderByRestaurantIdDateAndStatus)

	adminRoutes := app.Group("/api/admin")
	adminRoutes.Use(middleware.RequireAdminRole)

	adminRoutes.Get("/users", userInputAdapter.GetAllUsers)
	adminRoutes.Get("/owners", userInputAdapter.GetAllOwners)
	adminRoutes.Post("/user", userInputAdapter.CreateUser)
	adminRoutes.Get("/user/:userId", userInputAdapter.GetUserByUserId)
	adminRoutes.Put("/user/:userId", userInputAdapter.UpdateUser)
	adminRoutes.Delete("/user/:userId", userInputAdapter.DeleteUser)

	adminRoutes.Get("/restaurants", restaurantInputAdapter.GetAllRestaurants)
	adminRoutes.Patch("/restaurant/:restaurantId/status", restaurantInputAdapter.AdminUpdateRestaurantStatus)

	if err := app.Listen(":" + config.AppConfig.APIPort); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
