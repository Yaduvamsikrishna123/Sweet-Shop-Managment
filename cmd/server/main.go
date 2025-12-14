package main

import (
	"log"
	"os"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/database"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()

	// Auto Migrate
	database.DB.AutoMigrate(&models.User{}, &models.Sweet{})

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
