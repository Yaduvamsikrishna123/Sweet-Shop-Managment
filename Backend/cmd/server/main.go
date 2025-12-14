package main

import (
	"log"
	"os"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/database"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()

	// Auto Migrate
	database.DB.AutoMigrate(&models.User{}, &models.Sweet{}, &models.Transaction{})

	// Seed Data
	var count int64
	database.DB.Model(&models.Sweet{}).Count(&count)
	if count == 0 {
		sweets := []models.Sweet{
			{Name: "Gulab Jamun", Category: "Syrup Based", Price: 15.0, Quantity: 100},
			{Name: "Rasgulla", Category: "Syrup Based", Price: 20.0, Quantity: 50},
			{Name: "Kaju Katli", Category: "Dry Fruit", Price: 40.0, Quantity: 30},
			{Name: "Mysore Pak", Category: "Ghee Based", Price: 25.0, Quantity: 60},
			{Name: "Jalebi", Category: "Fried", Price: 10.0, Quantity: 80},
			{Name: "Laddu", Category: "Flour Based", Price: 12.0, Quantity: 120},
		}
		database.DB.Create(&sweets)
		log.Println("Seeded database with initial sweets")
	}

	// Seed Admin User
	var adminCount int64
	database.DB.Model(&models.User{}).Where("role = ?", "admin").Count(&adminCount)
	if adminCount == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := models.User{
			Username: "admin",
			Password: string(hashedPassword),
			Role:     "admin",
		}
		database.DB.Create(&admin)
		log.Println("Seeded database with default admin user")
	}

	r := gin.Default()

	// CORS Configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
