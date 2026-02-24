package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, error := NewUser("John Doe", "john_doe@mail.com", "123456")
	assert.Nil(t, error)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john_doe@mail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, error := NewUser("John Doe", "john_doe@mail.com", "123456")
	assert.Nil(t, error)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("wrong_password"))
	assert.NotEqual(t, "123456", user.Password)
}
