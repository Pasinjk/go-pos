package domain

import (
	"github.com/pasinjk/go-pos/internal/domain/model"
)

type UserRepository interface {
	Save(user model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (model.User, error)
}
