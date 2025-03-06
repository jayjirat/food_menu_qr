package main

import (
	"backend-food-menu-qr/config"
	"backend-food-menu-qr/core/domain"
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
	app := fiber.New()

	app.Post("/api/user/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})
	app.Listen(":" + config.AppConfig.APIPort)
}
