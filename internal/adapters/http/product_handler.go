package http

import (
	// "github.com/gofiber/fiber/v2"
	// "github.com/pasinjk/go-pos/internal/domain/model"
	// "github.com/pasinjk/go-pos/internal/domain/model/response"
	"github.com/pasinjk/go-pos/internal/usecase"
)

type HttpProductHandler struct {
	service usecase.ProductService
}

func NewHttpProductHandler(service usecase.ProductService) *HttpProductHandler {
	return &HttpProductHandler{service: service}
}

// func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
// 	var product model.Product
// 	if err := c.BodyParser(&product); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
// 	}
// 	savedProduct, err := h.service.CreateProduct(product)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON()

// }
