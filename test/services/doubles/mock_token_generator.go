package doubles

import (
	"github.com/rafaph/notte-auth/domain/entities"
	"github.com/rafaph/notte-auth/domain/services"
	"github.com/stretchr/testify/mock"
)

type MockTokenGenerator struct {
	mock.Mock
}

func (m *MockTokenGenerator) Generate(user *entities.User) (string, error) {
	args := m.Called(user)
	token := args.Get(0).(string)
	err := args.Error(1)
	return token, err
}

func NewMockTokenGenerator(token string, err error) services.TokenGenerator {
	tokenGenerator := new(MockTokenGenerator)
	tokenGenerator.On("Generate", mock.Anything).Return(token, err)
	return tokenGenerator
}
