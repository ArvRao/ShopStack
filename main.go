package main

import (
	"log"
	"os"

	"github.com/ArvRao/shopstack/api/models"
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
	// if err := database.InitDB(); err != nil {
	// 	log.Fatalf("Failed to initialize database: %v", err)
	// }
	// defer database.CloseDB()

	// Run migrations
	database.SyncDatabase()

	// test db connection
	testDatabaseConnection()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, ShopStack!")
	})

	// Add a new endpoint to test the database connection
	/* app.Get("/db-test", func(c *fiber.Ctx) error {
		var result int
		err := database.DB.QueryRow(context.Background(), "SELECT 1").Scan(&result)
		if err != nil {
			return c.Status(500).SendString("Database connection failed")
		}
		return c.SendString("Database connection successful")
	})
	*/
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}

// testDatabaseConnection performs a basic CRUD operation to verify the database setup
func testDatabaseConnection() {
	db, err := database.OpenDb()
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer database.CloseDb(db)

	// Define the email of the test user
	testEmail := "test2@example.com"

	// Step 0: Delete any existing user with the same email
	db.Where("email = ?", testEmail).Delete(&models.User{})
	log.Println("Existing test user (if any) deleted")

	// Step 1: Create a test user
	testUser := models.User{
		Name:         "Test User",
		Email:        testEmail,
		PasswordHash: "hashedpassword", // Normally, you'd hash this, but for testing, it's fine
	}

	// Insert the test user into the database
	if err := db.Create(&testUser).Error; err != nil {
		log.Fatalf("Failed to insert test user: %v", err)
	}
	log.Println("Test user created successfully")

	// Step 2: Retrieve the test user to verify it was inserted
	var retrievedUser models.User
	if err := db.First(&retrievedUser, "email = ?", testEmail).Error; err != nil {
		log.Fatalf("Failed to retrieve test user: %v", err)
	}
	log.Printf("Test user retrieved successfully: %+v\n", retrievedUser)

	// Step 3: Clean up (delete the test user)
	if err := db.Delete(&retrievedUser).Error; err != nil {
		log.Fatalf("Failed to delete test user: %v", err)
	}
	log.Println("Test user deleted successfully")

	log.Println("Database connection and basic CRUD test completed successfully!")
}
