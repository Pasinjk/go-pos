package model

import (
	"gorm.io/gorm"
)

type StockMovement struct {
	gorm.Model
	ProductID uint    `json:"product_id" validate:"required"`
	Product   Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity" validate:"required"`
	Reason    string  `json:"reason" validate:"required,oneof=restock sale return adjustment damage"`
	Notes     string  `json:"notes"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type StockRequest struct {
	ProductID uint   `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	Reason    string `json:"reason" validate:"required,oneof=restock sale return adjustment damage"`
	Notes     string `json:"notes"`
}
