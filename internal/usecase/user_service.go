package usecase

import (
	"errors"

	"github.com/pasinjk/go-pos/internal/domain"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (model.User, error)
}

type userServiceImpl struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user model.User) (model.User, error) {
	// Check required
	if user.Email == "" || user.Password == "" || user.Name == "" || user.Role == "" {
		return model.User{}, errors.New("Please check request data again")
	}

	// create hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, errors.New("error to hashed password")
	}

	user.Password = string(hashedPassword)

	// Save if all check is pass
	saveUser, err := s.repo.Save(user)
	if err != nil {
		return model.User{}, err
	}

	return saveUser, nil
}

func (s *userServiceImpl) GetAllUsers() ([]model.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userServiceImpl) GetUserByID(id uint) (model.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
