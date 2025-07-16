package response

import (
	"time"

	"github.com/pasinjk/go-pos/internal/domain/model"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type getUserIdResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
	}
}

func GetUserByIDResponse(user model.User) getUserIdResponse {
	return getUserIdResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
