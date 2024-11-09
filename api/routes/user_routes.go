package routes

import (
	"github.com/ArvRao/shopstack/api/controllers"
	"github.com/ArvRao/shopstack/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	// Registration route
	app.Post("/api/register", controllers.RegisterUser)
	// Login route
	app.Post("/api/login", controllers.LoginUser)

	// Protected routes (example: profile route)
	app.Get("/api/user/profile", middlewares.JWTMiddleware, controllers.GetUserProfile)
}
