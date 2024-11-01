package services

import (
	"WeatherProfile_Service/models"
	"WeatherProfile_Service/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}
