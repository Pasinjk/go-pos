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
	GetCategoryByID(id uint) (model.Category, error)
	DeleteCatagoryByID(id uint) error
}

type categoriesServiceImpl struct {
	repo domain.CategoriesRepository
}

func NewCategoriesService(repo domain.CategoriesRepository) CategoriesService {
	return &categoriesServiceImpl{repo: repo}
}

func (s *categoriesServiceImpl) CreateCategory(category model.Category) (model.Category, error) {
	if category.Name == "" {
		return model.Category{}, errors.New("category name is required")
	}

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

func (s *categoriesServiceImpl) GetCategoryByID(id uint) (model.Category, error) {
	if id == 0 {
		return model.Category{}, errors.New("category ID is required")
	}

	categories, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return model.Category{}, err
	}

	// for _, category := range categories {
	// 	if category.ID == id {
	// 		return category, nil
	// 	}
	// }

	return categories, nil
}

func (s *categoriesServiceImpl) DeleteCatagoryByID(id uint) error {
	if id == 0 {
		return errors.New("category ID is required")
	}
	err := s.repo.DeleteCatagoryByID(id)
	if err != nil {
		return err
	}
	return nil
}
