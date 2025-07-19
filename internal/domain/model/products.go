package model

import (
	"gorm.io/gorm"
	"time"
)

type Status string

const (
	InStock  Status = "in_stock"
	LowStock Status = "low_stock"
	OutStock Status = "out_of_stock"
)

type Product struct {
	gorm.Model
	Name          string  `json:"name" gorm:"not null" validate:"required,min=2,max=200"`
	Description   string  `json:"description"`
	Price         float64 `json:"price" gorm:"type:decimal(10,2);not null" validate:"required,min=0"`
	Cost          float64 `json:"cost" gorm:"type:decimal(10,2)" validate:"min=0"`
	CategoryId    uint    `gorm:"not null" json:"category_id"`
	Category      Category
	CategoryName  string `json:"category_name" gorm:"-"`
	StockQuantity uint   `json:"stock_quantity" gorm:"default:0" validate:"min=0"`
	MinStockLevel uint   `json:"min_stock_level" gorm:"default:0" validate:"min=0"`
}

type ProductRequest struct {
	Name          string  `json:"name" validate:"required,min=2,max=200"`
	Description   string  `json:"description"`
	Price         float64 `json:"price" validate:"required,min=0"`
	Cost          float64 `json:"cost" validate:"min=0"`
	CategoryID    uint    `json:"category_id" validate:"required"`
	StockQuantity uint    `json:"stock_quantity" validate:"min=0"`
	MinStockLevel uint    `json:"min_stock_level" validate:"min=0"`
}

type InventoryItem struct {
	ProductID     uint      `json:"product_id"`
	ProductName   string    `json:"product_name"`
	CurrentStock  uint      `json:"current_stock"`
	MinStockLevel uint      `json:"min_stock_level"`
	Status        Status    `json:"status"` // in_stock, low_stock, out_of_stock
	LastUpdated   time.Time `json:"last_updated"`
}
