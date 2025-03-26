package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// models "mh_api_go/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Database connection string
	dsn := "host=localhost user=root password=root dbname=mhdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
	log.Println("Database connected successfully!")
}
