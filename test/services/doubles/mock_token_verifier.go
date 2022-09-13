package doubles

import (
	"github.com/rafaph/notte-auth/domain/entities"
	"github.com/rafaph/notte-auth/domain/services"
	"github.com/stretchr/testify/mock"
)

type MockTokenVerifier struct {
	mock.Mock
	Returneduser *entities.User
}

func (m *MockTokenVerifier) Verify(token string) (*entities.User, error) {
	args := m.Called(token)
	user := args.Get(0).(*entities.User)
	err := args.Error(1)
	return user, err
}

func NewMockTokenVerifier(user *entities.User, err error) services.TokenVerifier {
	tokenVerifier := new(MockTokenVerifier)
	tokenVerifier.On("Verify", mock.Anything).Return(user, err)
	tokenVerifier.Returneduser = user
	return tokenVerifier
}
