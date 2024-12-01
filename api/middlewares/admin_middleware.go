package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// AdminMiddleware ensures that only admins can access certain routes
func AdminMiddleware(c *fiber.Ctx) error {
	// Extract the role from the JWT (stored in locals by JWT middleware)
	role := c.Locals("role")
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied. Admins only.",
		})
	}

	// Proceed to the next handler
	return c.Next()
}
