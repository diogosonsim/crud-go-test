package repository

import (
	"github.com/diogosonsim/crud-go-test/database"
	"github.com/diogosonsim/crud-go-test/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repository *UserRepository) GetUser(id uint) (*models.User, error) {
	var user models.User

	err := database.DB.Where("id = ?", id).
		First(&user).
		Error

	return &user, err
}

func (repository *UserRepository) GetUsers() (*[]models.User, error) {
	var users []models.User

	err := database.DB.Find(&users).Error

	return &users, err
}

func (repository *UserRepository) CreateUser(createUser *models.User) (*models.User, error) {
	user := models.User{
		Name:    createUser.Name,
		Age:     createUser.Age,
		Email:   createUser.Email,
		Address: createUser.Address,
	}
	user.SetPassword("12345")

	err := database.DB.Create(&user).Error

	return &user, err
}

func (repository *UserRepository) UpdateUser(updateUser *models.User) (*models.User, error) {
	user := models.User{}

	if err := database.DB.Where("id = ?", updateUser.Id).First(&user).Error; err != nil {
		return nil, err
	}

	err := database.DB.Updates(&updateUser).Error

	return updateUser, err
}

func (repository *UserRepository) DeleteUser(id uint) (string, error) {
	user := models.User{
		Id: uint(id),
	}

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return "", err
	}

	err := database.DB.Delete(&user).Error

	return "success", err
}
