package entity

import (
	"github.com/felipecaue-lb/goexpert/09-apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, erro := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if erro != nil {
		return nil, erro
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return error == nil
}
