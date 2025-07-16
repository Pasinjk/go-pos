package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string
	Price        float64
	Stock        uint
	CategoryId   uint
	Category     Category
	SalesItemId  uint
	SalesItem    SalesItem
	InventoryLog []InventoryLog
}
