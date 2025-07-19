package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name         string    `json:"name" gorm:"unique;not null" validate:"required,min=2,max=100"`
	Description  string    `json:"description"`
	ProductCount uint      `json:"product_count" gorm:"-"`
	Products     []Product `json:"products,omitempty"`
}

type AllCategoryDataResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ProductCount uint      `json:"product_count"`
	CreatedAt    time.Time `json:"created_at"`
}

func GetAllCategoryResponse(c Category) AllCategoryDataResponse {
	return AllCategoryDataResponse{
		ID:           c.ID,
		Name:         c.Name,
		Description:  c.Description,
		ProductCount: uint(len(c.Products)),
		CreatedAt:    c.CreatedAt,
	}
}

type CategoryDataResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products"`
	CreatedAt   time.Time `json:"created_at"`
}

func CategoryResponse(c Category) CategoryDataResponse {
	return CategoryDataResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Products:    c.Products,
		CreatedAt:   c.CreatedAt,
	}
}
