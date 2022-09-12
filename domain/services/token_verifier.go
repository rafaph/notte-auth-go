package services

import "github.com/rafaph/notte-auth/domain/entities"

type TokenVerifier interface {
	Verify(token string) (*entities.User, error)
}
