package controllers

import (
	"strings"
	"time"

	"github.com/ArvRao/shopstack/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// LogoutHandler handles the logout request
func LogoutHandler(c *fiber.Ctx) error {
	// Extract the token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse the token to get the expiration time
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	exp := int64(claims["exp"].(float64)) // Get expiration time from the token

	// Calculate the remaining validity period of the token
	expiration := time.Until(time.Unix(exp, 0))

	// Add the token to Redis with the remaining expiration time
	err = database.RedisClient.Set(database.Ctx, tokenString, "true", expiration).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to blacklist token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
