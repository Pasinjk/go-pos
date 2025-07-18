package usecase

import (
	"errors"

	"github.com/pasinjk/go-pos/internal/domain"
	"github.com/pasinjk/go-pos/internal/domain/model"
)

type CategoriesService interface {
	CreateCategory(category model.Category) (model.Category, error)
	GetAllCategories() ([]model.Category, error)
	UpdateCategory(category model.Category) (model.Category, error)
}

type categoriesServiceImpl struct {
	repo domain.CategoriesRepository
}

func NewCategoriesService(repo domain.CategoriesRepository) CategoriesService {
	return &categoriesServiceImpl{repo: repo}
}

func (s *categoriesServiceImpl) CreateCategory(category model.Category) (model.Category, error) {
	// Check required fields
	if category.Name == "" {
		return model.Category{}, errors.New("category name is required")
	}

	// Save the category using the repository
	savedCategory, err := s.repo.Save(category)
	if err != nil {
		return model.Category{}, err
	}

	return savedCategory, nil
}

func (s *categoriesServiceImpl) GetAllCategories() ([]model.Category, error) {
	categories, err := s.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *categoriesServiceImpl) UpdateCategory(category model.Category) (model.Category, error) {
	// Check required fields
	if category.ID == 0 || category.Name == "" {
		return model.Category{}, errors.New("category ID and name are required")
	}

	// Update the category using the repository
	updatedCategory, err := s.repo.UpdateCategory(category)
	if err != nil {
		return model.Category{}, err
	}

	return updatedCategory, nil
}
