package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/rafaph/notte-auth/infrastructure/clients"
	"github.com/rafaph/notte-auth/infrastructure/repositories"
	"github.com/rafaph/notte-auth/test/clients/doubles"
)

type HttpGetUserRepositoryTypes struct {
	ReturnedGetUserResponse *clients.GetUserResponse
	ReturnedError           error
	Client                  clients.UserClient
	Repository              *repositories.HttpGetUserRepository
}

func MakeHttpGetUserRepository(suts ...*HttpGetUserRepositoryTypes) *HttpGetUserRepositoryTypes {
	var sut *HttpGetUserRepositoryTypes = nil

	if len(suts) > 0 {
		sut = suts[0]
	}

	if sut == nil {
		sut = &HttpGetUserRepositoryTypes{}
	}

	sut.ReturnedGetUserResponse = &clients.GetUserResponse{
		Id: faker.UUIDHyphenated(),
	}
	sut.ReturnedError = nil

	client := doubles.NewMockUserClient(sut.ReturnedGetUserResponse, sut.ReturnedError)

	if sut.Client == nil {
		sut.Client = client
	}

	sut.Repository = repositories.NewHttpGetUserRepository(sut.Client)

	return sut
}
