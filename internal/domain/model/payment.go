package model

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	TransactionID   uint        `json:"transaction_id" validate:"required"`
	Transaction     Transaction `json:"transaction,omitempty" gorm:"foreignKey:TransactionID"`
	Amount          float64     `json:"amount" gorm:"type:decimal(10,2)" validate:"required,min=0"`
	PaymentMethod   string      `json:"payment_method" validate:"required,oneof=cash card digital_wallet check"`
	Status          string      `json:"status" gorm:"default:pending" validate:"oneof=pending completed failed cancelled"`
	ReferenceNumber string      `json:"reference_number"`
	ProcessedAt     *time.Time  `json:"processed_at"`
}

type PaymentRequest struct {
	TransactionID uint        `json:"transaction_id" validate:"required"`
	Amount        float64     `json:"amount" validate:"required,min=0"`
	PaymentMethod string      `json:"payment_method" validate:"required,oneof=cash card digital_wallet check"`
	CardDetails   *CardDetail `json:"card_details,omitempty"`
}

type CardDetail struct {
	CardNumber string `json:"card_number" validate:"required"`
	Expiry     string `json:"expiry" validate:"required"`
	CVV        string `json:"cvv" validate:"required"`
}

type PaymentResponse struct {
	PaymentID       string    `json:"payment_id"`
	Status          string    `json:"status"`
	TransactionID   uint      `json:"transaction_id"`
	Amount          float64   `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
	ProcessedAt     time.Time `json:"processed_at"`
	ReferenceNumber string    `json:"reference_number"`
}

type RefundRequest struct {
	PaymentID    string  `json:"payment_id" validate:"required"`
	RefundAmount float64 `json:"refund_amount" validate:"required,min=0"`
	Reason       string  `json:"reason" validate:"required"`
}

type RefundResponse struct {
	RefundID     string    `json:"refund_id"`
	PaymentID    string    `json:"payment_id"`
	RefundAmount float64   `json:"refund_amount"`
	Status       string    `json:"status"`
	ProcessedAt  time.Time `json:"processed_at"`
}
