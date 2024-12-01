package routes

import (
	"github.com/ArvRao/shopstack/api/controllers"
	"github.com/ArvRao/shopstack/api/middlewares"
	"github.com/ArvRao/shopstack/api/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoutes(app *fiber.App) {
	categoryService := &services.CategoryService{}
	categoryController := controllers.NewCategoryController(categoryService)

	// Admin-only routes
	app.Post("/api/categories", middlewares.JWTMiddleware, middlewares.AdminMiddleware, categoryController.CreateCategoryHandler)
	// app.Put("/api/categories/:id", middlewares.JWTMiddleware, middlewares.AdminMiddleware, categoryController.UpdateCategoryHandler)
	// app.Delete("/api/categories/:id", middlewares.JWTMiddleware, middlewares.AdminMiddleware, categoryController.DeleteCategoryHandler)

	// Public routes
	app.Get("/api/categories", categoryController.GetAllCategoriesHandler)
}
