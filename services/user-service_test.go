package services

import (
	"errors"
	"testing"

	"github.com/diogosonsim/crud-go-test/models"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetUserById(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.GetUser(1)
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "John Snow", user.Name)
		assert.Equal(t, "john.snow@winterfell.com", user.Email)
	}
}

func TestServiceGetUserByIdNotFound(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.GetUser(3)
	if assert.NotNil(t, err) && assert.Nil(t, user) {
		assert.Equal(t, "not found", err.Error())
	}
}

func TestServiceGetUserByEmail(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.GetUserByEmail("john.snow@winterfell.com")
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "John Snow", user.Name)
		assert.Equal(t, "john.snow@winterfell.com", user.Email)
	}
}

func TestServiceGetUserByEmailNotFound(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.GetUserByEmail("john.snow3@winterfell.com")
	if assert.NotNil(t, err) && assert.Nil(t, user) {
		assert.Equal(t, "not found", err.Error())
	}
}

func TestServiceGetUsers(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	users, err := s.GetUsers()
	if assert.Nil(t, err) && assert.NotNil(t, users) {
		assert.Equal(t, 2, len(*users))
	}
}

func TestServiceCreateUser(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	mockUser := models.User{
		Id:      3,
		Name:    "Teste",
		Email:   "teste@teste.com",
		Address: "Teste",
		Age:     10,
	}

	user, err := s.CreateUser(&mockUser)
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "Teste", user.Name)
		assert.Equal(t, "teste@teste.com", user.Email)
	}
}

func TestServiceCreateUserAlreadyExists(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	mockUser := models.User{
		Id:      3,
		Name:    "John Snow",
		Email:   "john.snow@winterfell.com",
		Address: "Teste",
		Age:     10,
	}

	user, err := s.CreateUser(&mockUser)
	if assert.NotNil(t, err) && assert.Nil(t, user) {
		assert.Equal(t, "Email already exist!", err.Error())
	}
}

func TestServiceUpdateUser(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	mockUser := models.User{
		Id:      2,
		Name:    "Bran Stark",
		Email:   "bran.stark@winterfell.com",
		Address: "Winterfell",
	}

	user, err := s.UpdateUser(&mockUser)
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "Bran Stark", user.Name)
		assert.Equal(t, "bran.stark@winterfell.com", user.Email)
	}
}

func TestServiceUpdateUserEmailAlreadyExists(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	mockUser := models.User{
		Id:      2,
		Name:    "Bran Stark",
		Email:   "john.snow@winterfell.com",
		Address: "Winterfell",
	}

	user, err := s.CreateUser(&mockUser)
	if assert.NotNil(t, err) && assert.Nil(t, user) {
		assert.Equal(t, "Email already exist!", err.Error())
	}
}

func TestServiceDeleteUser(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	msg, err := s.DeleteUser(1)
	if assert.Nil(t, err) && assert.NotNil(t, msg) {
		assert.Equal(t, "success", msg)
	}
}

func TestServiceDeleteUserNotFound(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	msg, err := s.DeleteUser(7)
	if assert.NotNil(t, err) && assert.NotNil(t, msg) {
		assert.Equal(t, "not found", err.Error())
	}
}

func (m *mockUserDAO) GetUser(id uint) (*models.User, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockUserDAO) GetUserByEmail(email string) (*models.User, error) {
	for _, record := range m.records {
		if record.Email == email {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockUserDAO) GetUsers() (*[]models.User, error) {
	return &m.records, nil
}

func (m *mockUserDAO) CreateUser(user *models.User) (*models.User, error) {
	for _, record := range m.records {
		if record.Email == user.Email {
			return nil, errors.New("Email already exist!")

		}
	}
	return user, nil
}

func (m *mockUserDAO) UpdateUser(user *models.User) (*models.User, error) {
	for _, record := range m.records {
		if record.Email == user.Email {
			return nil, errors.New("Email already exist!")

		}
	}
	return user, nil
}

func (m *mockUserDAO) DeleteUser(id uint) (string, error) {
	originalLength := len(m.records)

	index := 0
	for _, record := range m.records {
		if record.Id != id {
			m.records[index] = record
			index++
		}
	}
	newLength := len(m.records[:index])

	if originalLength == newLength {
		return "", errors.New("not found")
	} else {
		return "success", nil
	}
}

func newMockUserDAO() userRepository {
	return &mockUserDAO{
		records: []models.User{
			{Id: 1, Name: "John Snow", Email: "john.snow@winterfell.com", Address: "Winterfell"},
			{Id: 2, Name: "Ned Stark", Email: "ned.stark@winterfell.com", Address: "Winterfell"},
		},
	}
}

type mockUserDAO struct {
	records []models.User
	record  models.User
}
