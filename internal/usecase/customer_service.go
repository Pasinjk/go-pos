package usecase

import (
	"errors"

	"github.com/pasinjk/go-pos/internal/domain"
	"github.com/pasinjk/go-pos/internal/domain/model"
)

type CustomerService interface {
	CreateCustomer(customer model.Customer) (model.Customer, error)
}

type customerServiceImpl struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return &customerServiceImpl{repo: repo}
}

func (s *customerServiceImpl) CreateCustomer(customer model.Customer) (model.Customer, error) {
	if customer.Name == "" {
		return model.Customer{}, errors.New("customer name is required")
	}

	savedCustomer, err := s.repo.Save(customer)
	if err != nil {
		return model.Customer{}, err
	}

	return savedCustomer, nil
}
