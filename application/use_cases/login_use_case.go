package use_cases

import (
	"github.com/rafaph/notte-auth/domain/repositories"
	"github.com/rafaph/notte-auth/domain/services"
	"github.com/rafaph/notte-auth/lib/validator"
)

type LoginUseCaseInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type LoginUseCaseOutput struct {
	Token string
}

type LoginUseCase struct {
	getUserRepository repositories.GetUserRepository
	tokenGenerator    services.TokenGenerator
}

func (useCase *LoginUseCase) Execute(input LoginUseCaseInput) (*LoginUseCaseOutput, error) {
	err := validator.Validate(input)
	if err != nil {
		return nil, err
	}

	repositoryInput := repositories.GetUserInput{
		Email:    input.Email,
		Password: input.Password,
	}
	user, err := useCase.getUserRepository.GetUser(repositoryInput)
	if err != nil {
		return nil, err
	}

	token, err := useCase.tokenGenerator.Generate(user)
	if err != nil {
		return nil, err
	}

	return &LoginUseCaseOutput{Token: *token}, nil
}

func NewLoginUseCase(getUserRepository repositories.GetUserRepository, tokenGenerator services.TokenGenerator) *LoginUseCase {
	return &LoginUseCase{getUserRepository, tokenGenerator}
}
