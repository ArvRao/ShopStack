package routes

import (
	"github.com/ArvRao/shopstack/api/controllers"
	"github.com/ArvRao/shopstack/api/middlewares"
	"github.com/ArvRao/shopstack/api/services"
	"github.com/ArvRao/shopstack/database"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	// Initialize UserService and UserController
	userService := services.NewUserService(database.DB)
	userController := controllers.NewUserController(userService)

	// Registration route
	app.Post("/api/register", controllers.RegisterUser)
	// Login route
	app.Post("/api/login", controllers.LoginUser)
	// Protected routes
	app.Post("/api/logout", middlewares.JWT_Middleware, controllers.LogoutHandler)

	// Protected routes (example: profile route)
	app.Get("/api/user/profile", middlewares.JWTMiddleware, userController.GetUserProfileHandler)
	app.Put("/api/user/profile", middlewares.JWTMiddleware, userController.UpdateUserProfileHandler)
	app.Put("/api/user/change-password", middlewares.JWTMiddleware, userController.ChangePasswordHandler)
	// logout
}
