package database

import (
	"github.com/pasinjk/go-pos/internal/domain"
	// "github.com/pasinjk/go-pos/internal/domain/model"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) domain.ProductRepository {
	return &GormProductRepository{db: db}
}
