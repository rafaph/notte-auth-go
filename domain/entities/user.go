package entities

import "github.com/rafaph/notte-auth/lib/validator"

type User struct {
	Id string `validate:"required,uuid4"`
}

func NewUser(id string) (*User, error) {
	user := &User{
		Id: id,
	}

	err := validator.Validate(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
