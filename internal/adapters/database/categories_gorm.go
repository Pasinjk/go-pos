package database

import (
	"github.com/pasinjk/go-pos/internal/domain"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"gorm.io/gorm"
)

type GormCategoriesRepository struct {
	db *gorm.DB
}

func NewGormCategoriesRepository(db *gorm.DB) domain.CategoriesRepository {
	return &GormCategoriesRepository{db: db}
}

func (r *GormCategoriesRepository) Save(category model.Category) (model.Category, error) {
	if result := r.db.Create(&category); result.Error != nil {
		// Handle database errors
		return model.Category{}, result.Error
	}

	var savedCategory model.Category
	if err := r.db.First(&savedCategory, category.ID).Error; err != nil {
		// Handle database errors
		return model.Category{}, err
	}
	return savedCategory, nil
}

func (r *GormCategoriesRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *GormCategoriesRepository) UpdateCategory(category model.Category) (model.Category, error) {
	if result := r.db.Save(&category); result.Error != nil {
		return model.Category{}, result.Error
	}

	var updatedCategory model.Category
	if err := r.db.First(&updatedCategory, category.ID).Error; err != nil {
		return model.Category{}, err
	}

	return updatedCategory, nil
}
