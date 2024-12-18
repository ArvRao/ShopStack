package controllers

import (
	"github.com/ArvRao/shopstack/api/services"
	"github.com/gofiber/fiber/v2"
)

// CartController handles cart-related HTTP requests
type CartController struct {
	cartService *services.CartService
}

// NewCartController creates a new instance of CartController
func NewCartController(cartService *services.CartService) *CartController {
	return &CartController{cartService: cartService}
}

// AddToCartHandler handles POST /api/cart/add
func (cc *CartController) AddToCartHandler(c *fiber.Ctx) error {
	userIDInt := c.Locals("user_id").(int) // Get user_id from JWT middleware
	userID := uint(userIDInt)

	type request struct {
		ProductID uint `json:"product_id" validate:"required"`
		Quantity  int  `json:"quantity" validate:"required,min=1"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	cartItem, err := cc.cartService.AddToCart(userID, body.ProductID, body.Quantity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(cartItem)
}
