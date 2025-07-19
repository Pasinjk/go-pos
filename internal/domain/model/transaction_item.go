package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ReceiptNumber  string            `json:"receipt_number" gorm:"unique;not null"`
	CustomerID     uint              `json:"customer_id"`
	Customer       Customer          `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	CustomerName   string            `json:"customer_name" gorm:"-"`
	TotalAmount    float64           `json:"total_amount" gorm:"type:decimal(12,2);not null"`
	TaxAmount      float64           `json:"tax_amount" gorm:"type:decimal(10,2);default:0"`
	DiscountAmount float64           `json:"discount_amount" gorm:"type:decimal(10,2);default:0"`
	PaymentMethod  string            `json:"payment_method" validate:"required,oneof=cash card digital_wallet check"`
	Status         string            `json:"status" gorm:"default:pending" validate:"oneof=pending completed cancelled refunded"`
	Items          []TransactionItem `json:"items" gorm:"foreignKey:TransactionID"`
	Payments       []Payment         `json:"payments,omitempty" gorm:"foreignKey:TransactionID"`
	UserID         uint              `json:"user_id"`
	User           User              `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type TransactionItem struct {
	gorm.Model
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id" validate:"required"`
	Product       Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	ProductName   string  `json:"product_name" gorm:"-"`
	Quantity      int     `json:"quantity" validate:"required,min=1"`
	UnitPrice     float64 `json:"unit_price" gorm:"type:decimal(10,2)" validate:"required,min=0"`
	Subtotal      float64 `json:"subtotal" gorm:"type:decimal(10,2)"`
}

type TransactionRequest struct {
	CustomerID     *uint                    `json:"customer_id"`
	Items          []TransactionItemRequest `json:"items" validate:"required,min=1,dive"`
	PaymentMethod  string                   `json:"payment_method" validate:"required,oneof=cash card digital_wallet check"`
	DiscountAmount float64                  `json:"discount_amount" validate:"min=0"`
	TaxRate        float64                  `json:"tax_rate" validate:"min=0,max=1"`
}

type TransactionUpdateRequest struct {
	Status       string  `json:"status" validate:"oneof=pending completed cancelled refunded"`
	RefundAmount float64 `json:"refund_amount" validate:"min=0"`
	Reason       string  `json:"reason"`
}

type TransactionItemRequest struct {
	ProductID uint    `json:"product_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required,min=1"`
	UnitPrice float64 `json:"unit_price" validate:"required,min=0"`
}
