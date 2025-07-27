package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name           string        `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Email          string        `json:"email" gorm:"unique" validate:"email"`
	Phone          string        `json:"phone"`
	Address        string        `json:"address"`
	TotalPurchases float64       `json:"total_purchases" gorm:"type:decimal(12,2);default:0"`
	PurchaseCount  uint          `json:"purchase_count" gorm:"default:0"`
	LastPurchase   time.Time     `json:"last_purchase"`
	Transactions   []Transaction `json:"recent_transactions,omitempty" gorm:"foreignKey:CustomerID;limit:5"`
}

type CustomerResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	TotalPurchases float64   `json:"total_purchases"`
	LastPurchase   time.Time `json:"last_purchase"`
	CreateAt       time.Time `json:"create_at"`
}

func GetAllCustomerResponse(c Customer) CustomerResponse {
	return CustomerResponse{
		ID:             c.ID,
		Name:           c.Name,
		Email:          c.Email,
		Phone:          c.Phone,
		Address:        c.Address,
		TotalPurchases: c.TotalPurchases,
		LastPurchase:   c.LastPurchase,
		CreateAt:       c.CreatedAt,
	}
}

type GetCustomerResponseData struct {
	ID             uint          `json:"id"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Phone          string        `json:"phone"`
	Address        string        `json:"address"`
	TotalPurchases float64       `json:"total_purchases"`
	PurchaseCount  uint          `json:"purchase_count"`
	Transactions   []Transaction `json:"recent_transactions"`
	LastPurchase   time.Time     `json:"last_purchase"`
	CreateAt       time.Time     `json:"create_at"`
}

func GetCustomerResponse(c Customer) GetCustomerResponseData {
	return GetCustomerResponseData{
		ID:             c.ID,
		Name:           c.Name,
		Email:          c.Email,
		Phone:          c.Phone,
		Address:        c.Address,
		TotalPurchases: c.TotalPurchases,
		PurchaseCount:  c.PurchaseCount,
		LastPurchase:   c.LastPurchase,
		Transactions:   c.Transactions,
	}
}
