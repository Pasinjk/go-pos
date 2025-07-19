package domain

import (
	"github.com/pasinjk/go-pos/internal/domain/model"
)

type CategoriesRepository interface {
	Save(category model.Category) (model.Category, error)
	GetAllCategories() ([]model.Category, error)
	UpdateCategory(category model.Category) (model.Category, error)
	GetCategoryByID(id uint) (model.Category, error)
}
