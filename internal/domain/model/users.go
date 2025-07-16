package model

import (
	"gorm.io/gorm"
)

type Role string

const (
	AdminRole   Role = "admin"
	CashierRole Role = "cashier"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"not null" json:"name"`
	Email    string  `gorm:"unique;not null" json:"email"`
	Password string  `gorm:"not null" json:"password"`
	Role     Role    `gorm:"type:varchar(20);not null;default:'cashier'" json:"role"`
	Sales    []Sales `json:"sale"`
}
