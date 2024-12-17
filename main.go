package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ihksanghazi/API_delivery_service/database"
	"github.com/ihksanghazi/API_delivery_service/model/domain"
	"github.com/joho/godotenv"
)

func main() {
	// memuat file .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	// connect db
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbTimeZone := os.Getenv("DB_TIME_ZONE")

	database.ConnectDB(dbHost, dbUser, dbPassword, dbName, dbPort, dbTimeZone)
	database.DB.AutoMigrate(domain.User{}, domain.Address{}, domain.Shipment{}, domain.ShipmentDetail{}, domain.Pricing{}, domain.Payment{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello Dunia",
		})
	})

	app.Listen(":3000")
}