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
}
