package main

import (
	"go-talk/src/database"
	"go-talk/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDatabase()
	log.Println("Database connection successful!")

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8000")
}
