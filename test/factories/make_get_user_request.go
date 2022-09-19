package factories

import (
	"github.com/go-faker/faker/v4"
	. "github.com/rafaph/notte-auth/infrastructure/clients"
)

func MakeGetUserRequest() GetUserRequest {
	return GetUserRequest{Email: faker.Email(), Password: faker.Password()}
}
