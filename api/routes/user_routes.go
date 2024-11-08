package routes

import (
	"github.com/ArvRao/shopstack/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	// Registration route
	app.Post("/api/register", controllers.RegisterUser)
}
