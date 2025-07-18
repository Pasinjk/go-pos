package usecase

import (
	// "errors"

	"github.com/pasinjk/go-pos/internal/domain"
	// "github.com/pasinjk/go-pos/internal/domain/model"
)

type ProductService interface {
}

type productServiceImpl struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) ProductService {
	return &productServiceImpl{repo: repo}
}
