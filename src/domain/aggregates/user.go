package aggregates

import (
	"github.com/nanda03dev/go-ms-template/src/helper"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type User struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

// Business logic inside domain entities
func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

func NewUser(createUserDTO dto.CreateUserDTO) *User {
	return &User{
		UserID:   helper.Generate16DigitUUID(), // Generate unique ID (UUID or similar)
		Name:     createUserDTO.Name,
		Email:    createUserDTO.Email,
		Password: createUserDTO.Password,
	}

}

type UserRepository interface {
	Save(user *User) error
	FindById(id string) (*User, error)
}
