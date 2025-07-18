package response

import (
	"time"

	"github.com/pasinjk/go-pos/internal/domain/model"
)

type CategoriesResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewCategoriesResponse(category model.Category) CategoriesResponse {
	return CategoriesResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
	}
}
