package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/rafaph/notte-auth/domain/repositories"
)

func MakeGetUserInput() repositories.GetUserInput {
	return repositories.GetUserInput{
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}
