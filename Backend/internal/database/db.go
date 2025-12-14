package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// Use SQLite for local development to avoid auth issues
	DB, err = gorm.Open(sqlite.Open("sweetshop.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to SQLite database successfully")
}
