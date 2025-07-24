package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"github.com/pasinjk/go-pos/internal/usecase"
)

type HttpCustomerHandler struct {
	service usecase.CustomerService
}

func NewHttpCustomerHandler(service usecase.CustomerService) *HttpCustomerHandler {
	return &HttpCustomerHandler{service: service}
}

func (h *HttpCustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var customer model.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	saveCustomer, err := h.service.CreateCustomer(customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"name":    saveCustomer.Name,
		"email":   saveCustomer.Email,
		"phone":   saveCustomer.Phone,
		"address": saveCustomer.Address,
	})
}
