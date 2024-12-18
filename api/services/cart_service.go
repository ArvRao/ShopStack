package services

import (
	"errors"

	"github.com/ArvRao/shopstack/api/models"
	"github.com/ArvRao/shopstack/database"
)

// CartService handles cart-related operations
type CartService struct{}

// AddToCart adds a product to the user's cart or updates the quantity if it already exists
func (cs *CartService) AddToCart(userID, productID uint, quantity int) (*models.CartItem, error) {
	// Validate that the product exists and is in stock
	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		return nil, errors.New("product not found")
	}
	if product.Stock < quantity {
		return nil, errors.New("not enough stock available")
	}

	// Check if the product is already in the user's cart
	var cartItem models.CartItem
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error; err == nil {
		// Product is already in the cart; update the quantity
		newQuantity := cartItem.Quantity + quantity
		if product.Stock < newQuantity {
			return nil, errors.New("not enough stock available for the updated quantity")
		}
		cartItem.Quantity = newQuantity
		if err := database.DB.Save(&cartItem).Error; err != nil {
			return nil, errors.New("failed to update cart item")
		}
		return &cartItem, nil
	}

	// Product is not in the cart; create a new cart item
	cartItem = models.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}
	if err := database.DB.Create(&cartItem).Error; err != nil {
		return nil, errors.New("failed to add product to cart")
	}

	return &cartItem, nil
}
