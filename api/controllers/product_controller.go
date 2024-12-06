package controllers

import (
	"strconv"

	"github.com/ArvRao/shopstack/api/services"
	"github.com/gofiber/fiber/v2"
)

// ProductController handles product-related HTTP requests
type ProductController struct {
	productService *services.ProductService
}

// NewProductController creates a new instance of ProductController
func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) CreateProductHandler(c *fiber.Ctx) error {
	// Extract user_id and role from Locals
	userIDInt := c.Locals("user_id").(int)
	userID := uint(userIDInt)
	role := c.Locals("role").(string)

	// Allow only vendors or admins to create products
	if role != "vendor" && role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Only vendors or admins can create products",
		})
	}

	type request struct {
		CategoryID  uint    `json:"category_id" validate:"required"`
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" validate:"required"`
		Stock       int     `json:"stock" validate:"required"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	product, err := pc.productService.CreateProduct(userID, body.CategoryID, body.Name, body.Description, body.Price, body.Stock)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

// UpdateProductHandler handles PUT /api/products/:id (Vendor/Admin)
func (pc *ProductController) UpdateProductHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	type request struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err = pc.productService.UpdateProduct(uint(productID), userID, body.Name, body.Description, body.Price, body.Stock)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Product updated successfully"})
}

// DeleteProductHandler handles DELETE /api/products/:id (Vendor/Admin)
func (pc *ProductController) DeleteProductHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	err = pc.productService.DeleteProduct(uint(productID), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}

// GetAllProductsHandler handles GET /api/products (Public)
func (pc *ProductController) GetAllProductsHandler(c *fiber.Ctx) error {
	categoryIDParam := c.Query("category_id")
	var categoryID *uint
	if categoryIDParam != "" {
		id, err := strconv.Atoi(categoryIDParam)
		if err == nil {
			categoryID = uintPtr(uint(id))
		}
	}

	products, err := pc.productService.GetAllProducts(categoryID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

// GetProductHandler handles GET /api/products/:id (Public)
func (pc *ProductController) GetProductHandler(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	product, err := pc.productService.GetProduct(uint(productID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

// Helper to create a pointer to a uint
func uintPtr(i uint) *uint {
	return &i
}
