package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() {
	dsn := "host=localhost user=gorm password=gorm port=5432 sslmode=disable"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("log.Fatal: couldnt connect to database: ", err)
		panic("Could not connect to database:")
	}
}
