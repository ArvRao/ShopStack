package controllers

import (
	"strconv"

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

// UpdateCartItemHandler handles PUT /api/cart/update/:id
func (cc *CartController) UpdateCartItemHandler(c *fiber.Ctx) error {
	userIDInt := c.Locals("user_id").(int) // Get user_id from JWT middleware
	userID := uint(userIDInt)

	// Get the cart item ID from the URL params
	cartItemID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cart item ID"})
	}

	type request struct {
		NewQuantity int `json:"new_quantity" validate:"required"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	cartItem, err := cc.cartService.UpdateCartItem(userID, uint(cartItemID), body.NewQuantity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(cartItem)
}

// RemoveCartItemHandler handles DELETE /api/cart/remove/:id
func (cc *CartController) RemoveCartItemHandler(c *fiber.Ctx) error {
	userIDInt := c.Locals("user_id").(int) // Get user_id from JWT middleware
	userID := uint(userIDInt)

	// Get the cart item ID from the URL params
	cartItemID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cart item ID"})
	}

	// Call the service to remove the cart item
	err = cc.cartService.RemoveCartItem(userID, uint(cartItemID))
	if err != nil {
		if err.Error() == "cart item not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to remove cart item"})
	}

	return c.JSON(fiber.Map{"message": "Cart item removed successfully"})
}

// ViewCartHandler handles GET /api/cart
func (cc *CartController) ViewCartHandler(c *fiber.Ctx) error {
	userIDInt := c.Locals("user_id").(int) // Get user_id from JWT middleware
	userID := uint(userIDInt)

	// Call the service to retrieve the cart items
	cartItems, err := cc.cartService.ViewCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Calculate totals for the cart
	totalQuantity := 0
	totalPrice := 0.0
	for _, item := range cartItems {
		totalQuantity += item.Quantity
		totalPrice += float64(item.Quantity) * item.Product.Price
	}

	// Return cart details with totals
	return c.JSON(fiber.Map{
		"cart_items":    cartItems,
		"total_quantity": totalQuantity,
		"total_price":    totalPrice,
	})
}

// ClearCartHandler handles DELETE /api/cart/clear
func (cc *CartController) ClearCartHandler(c *fiber.Ctx) error {
	userIDInt := c.Locals("user_id").(int) // Get user_id from JWT middleware
	userID := uint(userIDInt)

	// Call the service to clear the cart
	err := cc.cartService.ClearCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to clear cart"})
	}

	return c.JSON(fiber.Map{"message": "Cart cleared successfully"})
}
