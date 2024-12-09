package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/API_delivery_service/database"
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

	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"testing": "ok",
		})
	})

	r.Run(":5000")
}
