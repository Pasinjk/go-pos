package database

import (
	"github.com/pasinjk/go-pos/internal/domain"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) domain.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user model.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	return nil
}
