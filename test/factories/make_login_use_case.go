package factories

import (
	"github.com/go-faker/faker/v4"
	. "github.com/rafaph/notte-auth/application/use_cases"
	"github.com/rafaph/notte-auth/domain/entities"
	"github.com/rafaph/notte-auth/domain/repositories"
	"github.com/rafaph/notte-auth/domain/services"
	. "github.com/rafaph/notte-auth/test/repositories/doubles"
	. "github.com/rafaph/notte-auth/test/services/doubles"
)

type LoginUseCaseTypes struct {
	ReturnedToken     string
	ReturnedUser      entities.User
	GetUserRepository repositories.GetUserRepository
	TokenGenerator    services.TokenGenerator
	UseCase           *LoginUseCase
}

func MakeLoginUseCase(suts ...*LoginUseCaseTypes) *LoginUseCaseTypes {
	var sut *LoginUseCaseTypes = nil

	if len(suts) > 0 {
		sut = suts[0]
	}

	if sut == nil {
		sut = &LoginUseCaseTypes{}
	}

	token := faker.UUIDHyphenated()
	user, _ := entities.NewUser(faker.UUIDHyphenated())
	getRepository := NewMockGetUserRepository(user, nil)
	tokenGenerator := NewMockTokenGenerator(token, nil)

	if sut.GetUserRepository == nil {
		sut.GetUserRepository = getRepository
	}

	if sut.TokenGenerator == nil {
		sut.TokenGenerator = tokenGenerator
	}

	sut.ReturnedUser = *user
	sut.ReturnedToken = token
	sut.UseCase = NewLoginUseCase(sut.GetUserRepository, sut.TokenGenerator)

	return sut
}
