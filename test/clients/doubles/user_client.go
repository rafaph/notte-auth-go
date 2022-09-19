package doubles

import (
	"github.com/rafaph/notte-auth/infrastructure/clients"
	"github.com/stretchr/testify/mock"
)

type MockUserClient struct {
	mock.Mock
}

func (m *MockUserClient) GetUser(request clients.GetUserRequest) (*clients.GetUserResponse, error) {
	args := m.Called(request)
	response := args.Get(0).(*clients.GetUserResponse)
	err := args.Error(1)

	return response, err
}

func NewMockUserClient(response *clients.GetUserResponse, err error) clients.UserClient {
	userClient := new(MockUserClient)
	userClient.On("GetUser", mock.Anything).Return(response, err)
	return userClient
}
