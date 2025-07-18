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
