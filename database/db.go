package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(host string, user string, password string, dbName string, port string, timeZone string) {
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=" + timeZone
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Failed to Connection")
	} else {
		fmt.Println("Success Connect Database")
	}
}
