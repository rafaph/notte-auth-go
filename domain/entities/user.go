package entities

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id string `validate:"required,uuid4"`
}

func (user *User) validate() error {
	validate := validator.New()
	return validate.Struct(user)
}

func NewUser(id string) (*User, error) {
	user := &User{
		Id: id,
	}

	err := user.validate()

	if err != nil {
		return nil, err
	}

	return user, nil
}
