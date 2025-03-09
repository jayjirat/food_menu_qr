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

	db.AutoMigrate(&domain.User{})
}

func main() {
	initDB()
	userRepo := repositoryAdapter.NewUserOutputAdapter(db)
	userUsecase := usecase.NewUserUseCase(userRepo)
	authenticationUsecase := usecase.NewAuthenticationUseCase(userRepo)
	userInputAdapter := httpAdapter.NewUserInputAdapter(userUsecase)
	authenticationInputAdapter := httpAdapter.NewAuthenticationAdapter(authenticationUsecase)
	app := fiber.New()

	app.Post("/api/register", authenticationInputAdapter.Register)
	app.Post("/api/login", authenticationInputAdapter.Login)
	app.Use(middleware.AuthenticateToken)
	app.Get("/api/users", userInputAdapter.GetAllUsers)
	app.Get("/api/user/:id", userInputAdapter.GetUserByID)
	app.Post("/api/user", userInputAdapter.CreateUser)
	app.Put("/api/user", userInputAdapter.UpdateUser)

	app.Listen(":" + config.AppConfig.APIPort)
}
