package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ArvRao/shopstack/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and assigns it to the global DB variable
func InitDB() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open a new GORM database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Assign the connection to the global DB variable
	DB = db
	log.Println("Successfully connected to the database")

	return nil
}

// SyncDatabase runs the auto-migrations for all models
func SyncDatabase() {
	if DB == nil {
		log.Fatal("Database is not initialized")
	}

	// Add your models here for migration
	err := DB.AutoMigrate(&models.User{}, &models.Address{}, &models.Order{}, &models.Product{}, &models.ProductImage{}, &models.Category{}, &models.CartItem{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	log.Println("Database migrated successfully")
}
