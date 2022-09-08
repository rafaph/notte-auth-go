package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validate(class interface{}) error {
	return validate.Struct(class)
}
