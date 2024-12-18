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

// UpdateCartItem updates the quantity of a product in the user's cart
func (cs *CartService) UpdateCartItem(userID, cartItemID uint, newQuantity int) (*models.CartItem, error) {
	if newQuantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	// Find the cart item
	var cartItem models.CartItem
	if err := database.DB.Where("id = ? AND user_id = ?", cartItemID, userID).First(&cartItem).Error; err != nil {
		return nil, errors.New("cart item not found")
	}

	// Validate stock availability
	var product models.Product
	if err := database.DB.First(&product, cartItem.ProductID).Error; err != nil {
		return nil, errors.New("product not found")
	}
	if product.Stock < newQuantity {
		return nil, errors.New("not enough stock available")
	}

	// Update the cart item's quantity
	cartItem.Quantity = newQuantity
	if err := database.DB.Save(&cartItem).Error; err != nil {
		return nil, errors.New("failed to update cart item")
	}

	return &cartItem, nil
}

// RemoveCartItem removes a product from the user's cart
func (cs *CartService) RemoveCartItem(userID, cartItemID uint) error {
	// Find the cart item and ensure it belongs to the user
	var cartItem models.CartItem
	if err := database.DB.Where("id = ? AND user_id = ?", cartItemID, userID).First(&cartItem).Error; err != nil {
		return errors.New("cart item not found")
	}

	// Delete the cart item
	if err := database.DB.Delete(&cartItem).Error; err != nil {
		return errors.New("failed to remove cart item")
	}

	return nil
}

// ViewCart retrieves all items in the user's cart
func (cs *CartService) ViewCart(userID uint) ([]models.CartItem, error) {
	var cartItems []models.CartItem

	// Use GORM Preload to fetch related product details
	if err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return nil, errors.New("failed to retrieve cart items")
	}

	return cartItems, nil
}

// ClearCart removes all items from the user's cart
func (cs *CartService) ClearCart(userID uint) error {
	// Delete all cart items where user_id matches the given userID
	if err := database.DB.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
		return errors.New("failed to clear cart")
	}

	return nil
}
