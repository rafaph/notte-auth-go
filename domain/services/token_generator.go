package services

import "github.com/rafaph/notte-auth/domain/entities"

type TokenGenerator interface {
	Generate(*entities.User) (*string, error)
}
