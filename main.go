package main

import (
	"go-talk/src/database"
	"go-talk/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("failed to connect to database", err)
	} else {
		log.Println("Database connection successful!")
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8000")
}
