package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/rafaph/notte-auth/domain/entities"
)

func MakeUser() *entities.User {
	user, _ := entities.NewUser(faker.UUIDHyphenated())

	return user
}
