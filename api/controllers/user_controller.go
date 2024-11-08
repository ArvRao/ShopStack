package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/ArvRao/shopstack/api/models"
	"github.com/ArvRao/shopstack/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// UserRegistrationRequest defines the JSON structure for registration request
type UserRegistrationRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone,omitempty"`
}

// RegisterUser handles user registration
func RegisterUser(c *fiber.Ctx) error {
	// Parse the request body
	var req UserRegistrationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Hash the password
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process password",
		})
	}

	// Create a new user instance
	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Phone:        req.Phone,
		Role:         models.RoleCustomer, // Default role for registered users
		Status:       models.StatusActive, // Default status for new users
	}

	// Save user to the database
	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		// Check if the error is a unique constraint violation
		if isUniqueConstraintError(err) {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"error": "Email already in use",
			})
		}

		log.Printf("Failed to create user: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Success response
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

// hashPassword hashes the password using bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// isUniqueConstraintError checks if the error is a unique constraint violation
func isUniqueConstraintError(err error) bool {
	// Check if the error is a GORM error caused by a unique constraint violation
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return true
	}
	return false
}
