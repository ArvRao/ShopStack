package main

import (
	"context"
	"log"
	"os"

	"github.com/ArvRao/shopstack/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // Adjust the path based on your directory structure
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	myVar := os.Getenv("DB_HOST")
	log.Println("DB_HOST:", myVar)

	// Initialize database connection
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, ShopStack!")
	})

	// Add a new endpoint to test the database connection
	app.Get("/db-test", func(c *fiber.Ctx) error {
		var result int
		err := database.DB.QueryRow(context.Background(), "SELECT 1").Scan(&result)
		if err != nil {
			return c.Status(500).SendString("Database connection failed")
		}
		return c.SendString("Database connection successful")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
