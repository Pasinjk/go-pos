package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"github.com/pasinjk/go-pos/internal/domain/model/response"
	"github.com/pasinjk/go-pos/internal/usecase"
)

type HttpUserHandler struct {
	service usecase.UserService
}

func NewHttpUserHandler(service usecase.UserService) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}

func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	savedUser, err := h.service.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(response.NewUserResponse(savedUser))
}

func (h *HttpUserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.NewUserResponse(user))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": userResponses})
}

func (h *HttpUserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user ID is required"})
	}
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.service.GetUserByID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response.GetUserByIDResponse(user))
}
