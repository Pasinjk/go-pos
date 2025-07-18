package domain

import (
	"github.com/pasinjk/go-pos/internal/domain/model"
)

type UserRepository interface {
	Save(user model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(id uint) error
}
