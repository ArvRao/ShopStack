package controllers

import (
	"strconv"

	"github.com/ArvRao/shopstack/api/services"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

// CreateCategoryHandler handles POST /api/categories
func (cc *CategoryController) CreateCategoryHandler(c *fiber.Ctx) error {
	type request struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	category, err := cc.categoryService.CreateCategory(body.Name, body.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// UpdateCategoryHandler handles PUT /api/categories/:id
func (cc *CategoryController) UpdateCategoryHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
	}

	type request struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err = cc.categoryService.UpdateCategory(uint(id), body.Name, body.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Category updated successfully"})
}

// DeleteCategoryHandler handles DELETE /api/categories/:id
func (cc *CategoryController) DeleteCategoryHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
	}

	err = cc.categoryService.DeleteCategory(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}

// GetAllCategoriesHandler handles GET /api/categories
func (cc *CategoryController) GetAllCategoriesHandler(c *fiber.Ctx) error {
	categories, err := cc.categoryService.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(categories)
}
