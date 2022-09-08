package doubles

import (
	"github.com/rafaph/notte-auth/domain/entities"
	"github.com/rafaph/notte-auth/domain/repositories"
	"github.com/stretchr/testify/mock"
)

type MockGetUserRepository struct {
	mock.Mock
}

func (m *MockGetUserRepository) GetUser(input repositories.GetUserInput) (*entities.User, error) {
	args := m.Called(input)
	user := args.Get(0).(*entities.User)
	err := args.Error(1)
	return user, err
}

func NewMockGetUserRepository(user *entities.User, err error) repositories.GetUserRepository {
	getRepository := new(MockGetUserRepository)
	getRepository.On("GetUser", mock.Anything).Return(user, err)
	return getRepository
}
