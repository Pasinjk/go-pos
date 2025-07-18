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
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(id uint) error // Optional, if you want to implement user deletion
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
		return model.User{}, errors.New("please check request data again")
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

func (s *userServiceImpl) UpdateUser(user model.User) (model.User, error) {
	if user.ID == 0 {
		return model.User{}, errors.New("user ID is required for update")
	}

	// Check if user exists
	existingUser, err := s.repo.GetUserByID(user.ID)
	if err != nil {
		return model.User{}, err
	}

	// Update fields
	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Role != "" {
		existingUser.Role = user.Role
	}
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return model.User{}, errors.New("error to hashed password")
		}
		existingUser.Password = string(hashedPassword)
	}

	// Save updated user
	updatedUser, err := s.repo.UpdateUser(existingUser)
	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

func (s *userServiceImpl) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("user ID is required for deletion")
	}

	// Check if user exists
	_, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	// Delete user
	if err := s.repo.DeleteUser(id); err != nil {
		return err
	}

	return nil
}
