package controllers

import (
	"github.com/ArvRao/shopstack/api/services"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

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
