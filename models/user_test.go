package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPasswordEncryptAndPasswordCompare(t *testing.T) {
	userMock := User{}
	passwordMock := "teste"

	userMock.SetPassword(passwordMock)
	assert.NotEqual(t, passwordMock, userMock.Password)

	assert.Nil(t, userMock.ComparePassword(passwordMock))
}
