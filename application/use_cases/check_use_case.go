package use_cases

import "github.com/rafaph/notte-auth/domain/services"

type CheckUseCaseOutput struct {
	Id string
}

type CheckUseCase struct {
	tokenVerifier services.TokenVerifier
}

func (useCase *CheckUseCase) Execute(token string) (*CheckUseCaseOutput, error) {
	user, err := useCase.tokenVerifier.Verify(token)

	if err != nil {
		return nil, err
	}

	return &CheckUseCaseOutput{Id: user.Id}, nil
}

func NewCheckUseCase(tokenVerifier services.TokenVerifier) *CheckUseCase {
	return &CheckUseCase{tokenVerifier}
}
