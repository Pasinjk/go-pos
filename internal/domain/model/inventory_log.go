package model

import (
	"gorm.io/gorm"
)

type InventoryLog struct {
	gorm.Model
	ProductId uint
	Product   Product
	Change    uint
	Note      string
}
