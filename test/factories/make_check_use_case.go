package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/rafaph/notte-auth/application/use_cases"
	"github.com/rafaph/notte-auth/domain/entities"
	"github.com/rafaph/notte-auth/domain/services"
	"github.com/rafaph/notte-auth/test/services/doubles"
)

type CheckUseCaseTypes struct {
	TokenVerifier services.TokenVerifier
	UseCase       *use_cases.CheckUseCase
}

func MakeCheckUseCase(suts ...*CheckUseCaseTypes) *CheckUseCaseTypes {
	var sut *CheckUseCaseTypes = nil

	if len(suts) > 0 {
		sut = suts[0]
	}

	if sut == nil {
		sut = &CheckUseCaseTypes{}
	}

	user, _ := entities.NewUser(faker.UUIDHyphenated())
	tokenVerifier := doubles.NewMockTokenVerifier(user, nil)
	if sut.TokenVerifier == nil {
		sut.TokenVerifier = tokenVerifier
	}

	sut.UseCase = use_cases.NewCheckUseCase(sut.TokenVerifier)

	return sut
}
