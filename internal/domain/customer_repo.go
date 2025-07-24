package domain

import (
	"github.com/pasinjk/go-pos/internal/domain/model"
)

type CustomerRepository interface {
	Save(customer model.Customer) (model.Customer, error)
}
