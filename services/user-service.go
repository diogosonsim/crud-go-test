package services

import "github.com/diogosonsim/crud-go-test/models"

type userRepository interface {
	GetUsers() (*[]models.User, error)
	GetUser(id uint) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uint) (string, error)
}

type UserService struct {
	repository userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.repository.GetUser(id)
}

func (s *UserService) GetUsers() (*[]models.User, error) {
	return s.repository.GetUsers()
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.repository.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return s.repository.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) (string, error) {
	return s.repository.DeleteUser(id)
}
