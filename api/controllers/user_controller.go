package controllers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ArvRao/shopstack/api/dto"
	"github.com/ArvRao/shopstack/api/models"
	"github.com/ArvRao/shopstack/api/services"
	"github.com/ArvRao/shopstack/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// UserController encapsulates dependencies for user-related handlers
type UserController struct {
	userService *services.UserService
}

// NewUserController creates a new instance of UserController
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

var userService = services.NewUserService(database.DB) // Inject database dependency

// UserRegistrationRequest defines the JSON structure for registration request
type UserRegistrationRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone,omitempty"`
}

// UserLoginRequest defines the JSON structure for login request
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
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

// LoginUser handles user authentication
func LoginUser(c *fiber.Ctx) error {
	var req UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Find the user by email
	var user models.User
	db := database.DB
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Generate JWT token
	token, err := generateJWT(user)
	if err != nil {
		log.Printf("Failed to generate JWT: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Return the JWT token
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

// generateJWT generates a JWT token for the given user
func generateJWT(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GetUserProfile retrieves the user profile for the logged-in user
func GetUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	email := c.Locals("email")
	role := c.Locals("role")

	// Send user data back as a response
	return c.JSON(fiber.Map{
		"user_id": userID,
		"email":   email,
		"role":    role,
	})
}

// GetUserProfileHandler handles GET /api/user/profile
func (uc *UserController) GetUserProfileHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	profile, err := uc.userService.GetUserProfile(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(profile)
}

// UpdateUserProfileHandler handles PUT /api/user/profile
func (uc *UserController) UpdateUserProfileHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	var updateData dto.UpdateUserProfileRequest
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := uc.userService.UpdateUserProfile(uint(userID), &updateData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}

// ChangePasswordHandler handles PUT /api/user/change-password
func (uc *UserController) ChangePasswordHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	var req dto.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.NewPassword == "" || req.CurrentPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Both current and new passwords are required"})
	}

	err := uc.userService.ChangePassword(uint(userID), req.CurrentPassword, req.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}
