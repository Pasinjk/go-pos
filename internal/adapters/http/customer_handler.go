package http

import (
	"strconv"

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

func (h *HttpCustomerHandler) GetAllCustomer(c *fiber.Ctx) error {
	customer, err := h.service.GetAllCustomer()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var customersResponse []model.CustomerResponse
	for _, customers := range customer {
		customersResponse = append(customersResponse, model.GetAllCustomerResponse(customers))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customer": customersResponse,
		"total":    len(customersResponse),
	})
}

func (h *HttpCustomerHandler) GetCustomerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "category ID is required"})
	}
	customerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	customer, err := h.service.GetCustomerByID(uint(customerID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(model.GetCustomerResponse(customer))
}
