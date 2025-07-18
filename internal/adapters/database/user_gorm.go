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

func (r *GormUserRepository) Save(user model.User) (model.User, error) {
	if result := r.db.Create(&user); result.Error != nil {
		// Handle database errors
		return model.User{}, result.Error
	}

	var SavedUser model.User
	if err := r.db.First(&SavedUser, user.ID).Error; err != nil {
		// Handle database errors
		return model.User{}, err
	}

	return SavedUser, nil
}

func (r *GormUserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) GetUserByID(id uint) (model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *GormUserRepository) UpdateUser(user model.User) (model.User, error) {
	if result := r.db.Save(&user); result.Error != nil {
		return model.User{}, result.Error
	}

	var updatedUser model.User
	if err := r.db.First(&updatedUser, user.ID).Error; err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	if result := r.db.Delete(&model.User{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
