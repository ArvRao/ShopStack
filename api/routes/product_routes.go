package routes

import (
	"github.com/ArvRao/shopstack/api/controllers"
	"github.com/ArvRao/shopstack/api/middlewares"
	"github.com/ArvRao/shopstack/api/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App) {
	productService := &services.ProductService{}
	productController := controllers.NewProductController(productService)

	// Admin/Vendor routes
	app.Post("/api/products", middlewares.JWTMiddleware, productController.CreateProductHandler)
	app.Put("/api/products/:id", middlewares.JWTMiddleware, productController.UpdateProductHandler)
	app.Delete("/api/products/:id", middlewares.JWTMiddleware, productController.DeleteProductHandler)

	// Public routes
	app.Get("/api/products", productController.GetAllProductsHandler)
	app.Get("/api/products/:id", productController.GetProductHandler)
}
