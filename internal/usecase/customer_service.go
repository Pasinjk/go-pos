package usecase

import (
	"errors"

	"github.com/pasinjk/go-pos/internal/domain"
	"github.com/pasinjk/go-pos/internal/domain/model"
)

type CustomerService interface {
	CreateCustomer(customer model.Customer) (model.Customer, error)
	GetAllCustomer() ([]model.Customer, error)
	GetCustomerByID(id uint) (model.Customer, error)
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

func (s *customerServiceImpl) GetAllCustomer() ([]model.Customer, error) {
	customer, err := s.repo.GetAllCustomer()
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *customerServiceImpl) GetCustomerByID(id uint) (model.Customer, error) {
	if id == 0 {
		return model.Customer{}, errors.New("category ID is required")
	}
	customer, err := s.repo.GetCustomerByID(id)
	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}
