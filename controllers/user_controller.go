package services

import (
	"WeatherProfile_Service/models"
	"WeatherProfile_Service/repositories"
	"log"

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
		log.Printf("Error hashing password: %v", err)
		return err
	}
	user.Password = string(hashedPassword)
	err = s.repo.Create(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
	}
	return err
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		log.Printf("Error finding user by email (%s): %v", email, err)
		return nil, err
	}
	log.Printf("Found user: %+v", user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("Error comparing password for user (%s): %v", email, err)
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		log.Printf("Error finding user by ID (%d): %v", id, err)
	}
	return user, err
}
