package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/rafaph/notte-auth/application/use_cases"
)

func MakeLoginUseCaseInput() use_cases.LoginUseCaseInput {
	return use_cases.LoginUseCaseInput{
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}
