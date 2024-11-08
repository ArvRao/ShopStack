package main

import (
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

	// Run migrations
	database.SyncDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, ShopStack!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
