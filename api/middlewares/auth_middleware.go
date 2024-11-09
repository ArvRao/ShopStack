package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JWTMiddleware is a middleware function to protect routes using JWT
func JWTMiddleware(c *fiber.Ctx) error {
	// Get the JWT from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	}

	// Split the Authorization header to get the token part
	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token method conforms to HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}

		// Return the secret signing key
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// If there was an error in parsing or the token is invalid, return an error
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired JWT",
		})
	}

	// Get the claims from the token (assuming claims are of type jwt.MapClaims)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Failed to parse JWT claims",
		})
	}

	// Attach user information from claims to the context
	c.Locals("user_id", claims["user_id"])
	c.Locals("email", claims["email"])
	c.Locals("role", claims["role"])

	// Proceed to the next handler
	return c.Next()
}
