package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Address  string `json:"address"`
}

type UserRegister struct {
	Name            string `json:"name"`
	Age             int    `json:"age"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Address         string `json:"address"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Function to generate hashed password
func (user *User) SetPassword(value string) {
	password, _ := bcrypt.GenerateFromPassword([]byte(value), 14)
	user.Password = password
}

// Function to compare user password
func (user *User) ComparePassword(value string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(value))
}

func (user *User) IsValid() (errs []string) {
	if user.Id == 0 {
		errs = append(errs, "The id is required!")
	}
	if user.Name == "" {
		errs = append(errs, "The name is required!")
	}
	if len(user.Name) < 2 || len(user.Name) > 40 {
		errs = append(errs, "The name field must be between 2-40 chars!")
	}
	if user.Email == "" {
		errs = append(errs, "The email field is required!")
	}

	return errs
}

func (user *UserRegister) IsValidRegister() (errs []string) {
	if user.Name == "" {
		errs = append(errs, "The name is required!")
	}
	if len(user.Name) < 2 || len(user.Name) > 40 {
		errs = append(errs, "The name field must be between 2-40 chars!")
	}
	if user.Email == "" {
		errs = append(errs, "The email field is required!")
	}
	if user.Password == "" {
		errs = append(errs, "The password is required!")
	}
	if user.ConfirmPassword == "" {
		errs = append(errs, "The confirm_password is required!")
	}

	if user.Password != user.ConfirmPassword {
		errs = append(errs, "Passwords do not match")
	}

	return errs
}

func (user *LoginRequest) IsValidLogin() (errs []string) {
	if user.Email == "" {
		errs = append(errs, "The email field is required!")
	}
	if user.Password == "" {
		errs = append(errs, "The password is required!")
	}
	return errs
}
