package database

import (
	"github.com/pasinjk/go-pos/internal/domain"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &GormCustomerRepository{db: db}
}

func (r *GormCustomerRepository) Save(customer model.Customer) (model.Customer, error) {
	if result := r.db.Create(&customer); result.Error != nil {
		return model.Customer{}, result.Error
	}

	var savedCustomer model.Customer
	if err := r.db.First(&savedCustomer, customer.ID).Error; err != nil {
		return model.Customer{}, err
	}

	return savedCustomer, nil
}

func (r *GormCustomerRepository) GetAllCustomer() ([]model.Customer, error) {
	var customers []model.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *GormCustomerRepository) GetCustomerByID(id uint) (model.Customer, error) {
	var customer model.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}
