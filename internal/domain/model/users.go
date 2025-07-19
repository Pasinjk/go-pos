package model

import (
	"gorm.io/gorm"
	"time"
)

type Role string

const (
	AdminRole   Role = "admin"
	CashierRole Role = "cashier"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name" validate:"required,min=3,max=50"`
	Email    string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"omitempty,min=6"`
	Role     Role   `gorm:"type:varchar(20);not null;default:'cashier'" json:"role" validate:"required,oneof=admin cashier"`
}

type UserDataResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func UserResponse(user User) UserDataResponse {
	return UserDataResponse{
		ID:        user.ID,
		Username:  user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}
