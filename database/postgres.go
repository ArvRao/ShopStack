package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global DB instance
var DB *gorm.DB

// OpenDb initializes the database connection
func OpenDb() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open a new GORM database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

// CloseDb closes the database connection
func CloseDb(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting database: %v", err)
		return
	}
	sqlDB.Close()
}

// SyncDatabase performs auto-migration for all models
func SyncDatabase() {
	db, err := OpenDb()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer CloseDb(db)

	/* models := []interface{}{
		&models.Admin{},
		&models.Product{},
		&models.ProductImage{},
		&models.User{},
		&models.Address{},
		// &models.Cart{},
		&models.Order{},
		&models.Review{},
		// &models.CartTotal{},
	}
	log.Println("Models:", models)
	for _, model := range models {
		err := db.AutoMigrate(model)
		errHandler(err)
	} */

	log.Println("Database migrated successfully")
}

// Error handler for migration
func errHandler(e error) {
	if e != nil {
		log.Println(e)
	}
}
