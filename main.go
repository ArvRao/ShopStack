package main

import (
	"log"
	"os"

	"github.com/ArvRao/shopstack/api/routes"
	"github.com/ArvRao/shopstack/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database connection
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run database migrations
	database.SyncDatabase()

	// Initialize the Fiber app
	app := fiber.New()

	// Define root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the ShopStack!",
		})
	})

	// Register routes
	routes.RegisterUserRoutes(app)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
