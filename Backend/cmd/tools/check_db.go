package main

import (
	"fmt"
	"log"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sweetshop.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	var sweets []models.Sweet
	result := db.Find(&sweets)
	if result.Error != nil {
		log.Fatal("Failed to fetch sweets:", result.Error)
	}

	fmt.Printf("Found %d sweets:\n", len(sweets))
	for _, s := range sweets {
		fmt.Printf("- %d: %s (Qty: %d)\n", s.ID, s.Name, s.Quantity)
	}
}
