package services

import (
	"errors"

	"github.com/ArvRao/shopstack/api/models"
	"github.com/ArvRao/shopstack/database"
)

// CategoryService handles business logic related to categories
type CategoryService struct{}

// CreateCategory creates a new category
func (cs *CategoryService) CreateCategory(name, description string) (*models.Category, error) {
	category := &models.Category{
		Name:        name,
		Description: description,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		return nil, errors.New("failed to create category")
	}
	return category, nil
}

// UpdateCategory updates an existing category
func (cs *CategoryService) UpdateCategory(id uint, name, description string) error {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return errors.New("category not found")
	}

	category.Name = name
	category.Description = description

	if err := database.DB.Save(&category).Error; err != nil {
		return errors.New("failed to update category")
	}
	return nil
}

// DeleteCategory soft-deletes a category
func (cs *CategoryService) DeleteCategory(id uint) error {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return errors.New("category not found")
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return errors.New("failed to delete category")
	}
	return nil
}

// GetAllCategories retrieves all active categories
func (cs *CategoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Where("deleted_at IS NULL").Find(&categories).Error; err != nil {
		return nil, errors.New("failed to retrieve categories")
	}
	return categories, nil
}
