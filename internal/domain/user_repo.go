package domain

import (
	"github.com/pasinjk/go-pos/internal/domain/model"
)

type UserRepository interface {
	Save(user model.User) error
}
