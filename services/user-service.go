package services

import "github.com/diogosonsim/crud-go-test/models"

type userRepository interface {
	GetUsers() (*[]models.User, error)
	GetUser(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uint) (string, error)
}

// UserService struct
type UserService struct {
	repository userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

// Function to get user by id
func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.repository.GetUser(id)
}

// Function to get user by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repository.GetUserByEmail(email)
}

// Function to get all users
func (s *UserService) GetUsers() (*[]models.User, error) {
	return s.repository.GetUsers()
}

// Function to create a new user
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.repository.CreateUser(user)
}

// Function to update user values
func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return s.repository.UpdateUser(user)
}

// Function to delete user by id
func (s *UserService) DeleteUser(id uint) (string, error) {
	return s.repository.DeleteUser(id)
}
