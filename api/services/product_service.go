package services

import (
	"errors"

	"github.com/ArvRao/shopstack/api/models"
	"github.com/ArvRao/shopstack/database"
)

// ProductService handles business logic related to products
type ProductService struct{}

// CreateProduct creates a new product
func (ps *ProductService) CreateProduct(userID, categoryID uint, name, description string, price float64, stock int) (*models.Product, error) {
	product := &models.Product{
		UserID:      userID,
		CategoryID:  categoryID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return nil, errors.New("failed to create product")
	}
	return product, nil
}

// UpdateProduct updates an existing product
func (ps *ProductService) UpdateProduct(productID, userID uint, name, description string, price float64, stock int) error {
	var product models.Product
	if err := database.DB.Where("id = ? AND user_id = ?", productID, userID).First(&product).Error; err != nil {
		return errors.New("product not found or unauthorized")
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Stock = stock

	if err := database.DB.Save(&product).Error; err != nil {
		return errors.New("failed to update product")
	}
	return nil
}

// DeleteProduct soft-deletes a product
func (ps *ProductService) DeleteProduct(productID, userID uint) error {
	var product models.Product
	if err := database.DB.Where("id = ? AND user_id = ?", productID, userID).First(&product).Error; err != nil {
		return errors.New("product not found or unauthorized")
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}

// GetAllProducts retrieves all active products with optional filters
func (ps *ProductService) GetAllProducts(categoryID *uint) ([]models.Product, error) {
	var products []models.Product

	query := database.DB.Where("deleted_at IS NULL")
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, errors.New("failed to retrieve products")
	}
	return products, nil
}

// GetProduct retrieves a single product by ID
func (ps *ProductService) GetProduct(productID uint) (*models.Product, error) {
	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil
}
