package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/rafaph/notte-auth/infrastructure/repositories/http"
)

func MakeGetUserRequest() http.GetUserRequest {
	return http.GetUserRequest{Email: faker.Email(), Password: faker.Password()}
}
