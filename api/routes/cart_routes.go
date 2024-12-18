package routes

import (
	"github.com/ArvRao/shopstack/api/controllers"
	"github.com/ArvRao/shopstack/api/middlewares"
	"github.com/ArvRao/shopstack/api/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterCartRoutes(app *fiber.App) {
	cartService := &services.CartService{}
	cartController := controllers.NewCartController(cartService)

	// Protected cart routes
	app.Post("/api/cart/add", middlewares.JWTMiddleware, cartController.AddToCartHandler)
	app.Put("/api/cart/update/:id", middlewares.JWTMiddleware, cartController.UpdateCartItemHandler)
	app.Delete("/api/cart/remove/:id", middlewares.JWTMiddleware, cartController.RemoveCartItemHandler)
	app.Get("/api/cart", middlewares.JWTMiddleware, cartController.ViewCartHandler)
	app.Delete("/api/cart/clear", middlewares.JWTMiddleware, cartController.ClearCartHandler)
}
