package services

import (
	"crypto/ed25519"
	"github.com/kataras/jwt"
	"github.com/rafaph/notte-auth/config"
	"github.com/rafaph/notte-auth/domain/entities"
	"time"
)

type JwtTokenGenerator struct {
	config *config.JwtConfig
}

func (j *JwtTokenGenerator) getMaxAge() jwt.SignOptionFunc {
	return jwt.MaxAge(time.Duration(j.config.ExpirationInMinutes) * time.Minute)
}

func (j *JwtTokenGenerator) getPrivateKey() ed25519.PrivateKey {
	privateKey, _ := jwt.ParsePrivateKeyEdDSA([]byte(j.config.PrivateKey))

	return privateKey
}

func (j *JwtTokenGenerator) getClaims(user *entities.User) jwt.Map {
	return jwt.Map{"user_id": user.Id}
}

func (j *JwtTokenGenerator) Generate(user *entities.User) (string, error) {
	claims := j.getClaims(user)
	maxAge := j.getMaxAge()
	privateKey := j.getPrivateKey()

	tokenBytes, err := jwt.Sign(jwt.EdDSA, privateKey, claims, maxAge)

	if err != nil {
		return "", err
	}

	return jwt.BytesToString(tokenBytes), nil
}

func NewJwtTokenGenerator(jwtConfig *config.JwtConfig) *JwtTokenGenerator {
	return &JwtTokenGenerator{
		config: jwtConfig,
	}
}
