package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Email string
	Phone string
	Role  Role `gorm:"type:varchar(20);not null;default:'cashier'" json:"role"`
	Sales []Sales
}
