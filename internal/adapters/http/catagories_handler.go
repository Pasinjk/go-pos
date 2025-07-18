package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"github.com/pasinjk/go-pos/internal/domain/model/response"
	"github.com/pasinjk/go-pos/internal/usecase"
)

type HttpCategoriesHandler struct {
	service usecase.CategoriesService
}

func NewHttpCategoriesHandler(service usecase.CategoriesService) *HttpCategoriesHandler {
	return &HttpCategoriesHandler{service: service}
}

func (h *HttpCategoriesHandler) CreateCategory(c *fiber.Ctx) error {
	var category model.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	savedCategory, err := h.service.CreateCategory(category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"name":        savedCategory.Name,
		"description": savedCategory.Description,
	})
}

// TODO Count the product in each category and return each category with its product count
func (h *HttpCategoriesHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var categoryResponses []response.CategoriesResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, response.NewCategoriesResponse(category))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"categories": categoryResponses,
	})
}

func (h *HttpCategoriesHandler) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "category ID is required"})
	}
	categoryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var category model.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	category.ID = uint(categoryID)
	updatedCategory, err := h.service.UpdateCategory(category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"name":        updatedCategory.Name,
		"description": updatedCategory.Description})
}
