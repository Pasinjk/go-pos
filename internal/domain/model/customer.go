package model

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	Name           string        `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Email          string        `json:"email" gorm:"unique" validate:"email"`
	Phone          string        `json:"phone"`
	Address        string        `json:"address"`
	TotalPurchases float64       `json:"total_purchases" gorm:"type:decimal(12,2);default:0"`
	PurchaseCount  uint          `json:"purchase_count" gorm:"default:0"`
	LastPurchase   *time.Time    `json:"last_purchase"`
	Transactions   []Transaction `json:"recent_transactions,omitempty" gorm:"foreignKey:CustomerID;limit:5"`
}

type CustomerRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=100"`
	Email   string `json:"email" validate:"omitempty,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
