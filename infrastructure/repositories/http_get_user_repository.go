package repositories

import (
	"github.com/rafaph/notte-auth/domain/entities"
	"github.com/rafaph/notte-auth/domain/repositories"
	"github.com/rafaph/notte-auth/infrastructure/clients"
)

type HttpGetUserRepository struct {
	client clients.UserClient
}

func (h *HttpGetUserRepository) GetUser(input repositories.GetUserInput) (*entities.User, error) {
	response, err := h.client.GetUser(clients.GetUserRequest{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		return nil, err
	}

	user, err := entities.NewUser(response.Id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewHttpGetUserRepository(client clients.UserClient) *HttpGetUserRepository {
	return &HttpGetUserRepository{client}
}
