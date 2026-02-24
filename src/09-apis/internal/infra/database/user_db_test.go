package database

import (
	"testing"

	"github.com/felipecaue-lb/goexpert/09-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, error := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if error != nil {
		t.Error(error)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john_doe@mail.com", "123456")
	userDB := NewUser(db)

	error = userDB.Create(user)
	assert.Nil(t, error)

	var userFound entity.User
	error = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, error)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, error := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if error != nil {
		t.Error(error)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john_doe@mail.com", "123456")
	userDB := NewUser(db)

	error = userDB.Create(user)
	assert.Nil(t, error)

	userFound, error := userDB.FindByEmail(user.Email)
	assert.Nil(t, error)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
