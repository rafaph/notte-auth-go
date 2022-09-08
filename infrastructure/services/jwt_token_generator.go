package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/rafaph/notte-auth/config"
	"github.com/rafaph/notte-auth/domain/entities"
	"time"
)

type JwtTokenGenerator struct {
	config *config.JwtConfig
}

func (j *JwtTokenGenerator) getExpiration() time.Time {
	minutes := time.Duration(j.config.ExpirationInMinutes)

	return time.Now().Add(minutes * time.Minute)
}

func (j *JwtTokenGenerator) getSecret() []byte {
	return []byte(j.config.Secret)
}

func (j *JwtTokenGenerator) Generate(user *entities.User) (*string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = j.getExpiration()
	claims["user_id"] = user.Id

	secret := j.getSecret()
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return nil, err
	}

	return &tokenString, err
}

func NewJwtTokenGenerator(jwtConfig *config.JwtConfig) *JwtTokenGenerator {
	return &JwtTokenGenerator{
		config: jwtConfig,
	}
}
