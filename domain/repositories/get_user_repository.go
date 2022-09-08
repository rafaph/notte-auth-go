package repositories

import "github.com/rafaph/notte-auth/domain/entities"

type GetUserInput struct {
	Email    string
	Password string
}

type GetUserRepository interface {
	GetUser(GetUserInput) (*entities.User, error)
}
