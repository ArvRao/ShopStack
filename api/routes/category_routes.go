package routes

import (
	"github.com/ArvRao/shopstack/api/controllers"
	"github.com/ArvRao/shopstack/api/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoutes(app *fiber.App) {
	categoryService := &services.CategoryService{}
	categoryController := controllers.NewCategoryController(categoryService)

	app.Post("/api/categories", categoryController.CreateCategoryHandler)
}
