package model

import (
	"gorm.io/gorm"
)

type Sales struct {
	gorm.Model
	UserId        uint
	User          User
	CustomerId    uint
	Customer      Customer
	Total         float64
	PaymentMethod string
	Paid          bool
	SalesItem     []SalesItem
}

type SalesItem struct {
	gorm.Model
	SalesId  uint
	Sales    Sales
	Product  []Product
	Quantity uint
	Price    uint
	Subtotal uint
}
