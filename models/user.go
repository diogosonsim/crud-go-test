package models

import "golang.org/x/crypto/bcrypt"

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

func (user *User) SetPassword(value string) {
	password, _ := bcrypt.GenerateFromPassword([]byte(value), 14)
	user.Password = password
}

func (user *User) ComparePassword(value string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(value))
}
