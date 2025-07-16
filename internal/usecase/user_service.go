package usecase

import (
	"errors"

	"github.com/pasinjk/go-pos/internal/domain"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user model.User) error
}

type userServiceImpl struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user model.User) error {
	// Check required
	if user.Email == "" || user.Password == "" || user.Name == "" || user.Role == "" {
		return errors.New("Please check request data again")
	}

	// create hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error to hashed password")
	}

	user.Password = string(hashedPassword)

	// Save if all check is pass
	if err := s.repo.Save(user); err != nil {
		return err
	}
	return nil
}
