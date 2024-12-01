package middlewares

import (
	"os"
	"strings"

	"github.com/ArvRao/shopstack/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JWTMiddleware validates JWT tokens and checks for blacklisting
func JWT_Middleware(c *fiber.Ctx) error {
	// Ensure RedisClient is initialized
	if database.RedisClient == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Redis is not initialized",
		})
	}

	// Extract the token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Check if the token is blacklisted in Redis
	isBlacklisted, err := database.RedisClient.Get(database.Ctx, tokenString).Result()
	if err == nil && isBlacklisted == "true" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token is blacklisted. Please log in again.",
		})
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Extract claims and set them in Locals for subsequent handlers
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])
	c.Locals("email", claims["email"])

	return c.Next()
}
