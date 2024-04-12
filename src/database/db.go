package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable", dbUser, dbPassword, dbName)
	log.Println("dsn: ", dsn)

	maxAttempts := 5

	for attempt := 0; attempt < maxAttempts; attempt++ {

		log.Println(fmt.Sprintf("Database connection attempt %d\n", attempt))

		_, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Database connection successful")
			return nil
		} else {
			log.Fatal(fmt.Sprintf("attempt %d failed\n", attempt))
		}

		time.Sleep(10 * time.Second)
	}

	log.Fatal(fmt.Sprintf("failed to connect to database after %d attempts\n", maxAttempts))

	return err
}
