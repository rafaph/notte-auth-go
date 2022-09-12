package services

import (
	"crypto/ed25519"
	"fmt"
	"github.com/kataras/jwt"
	"github.com/rafaph/notte-auth/config"
	"github.com/rafaph/notte-auth/domain/entities"
)

type JwtTokenVerifier struct {
	config *config.JwtConfig
}

func (j *JwtTokenVerifier) getPublicKey() ed25519.PublicKey {
	publicKey, _ := jwt.ParsePublicKeyEdDSA([]byte(j.config.PublicKey))

	return publicKey
}

func (j *JwtTokenVerifier) Verify(token string) (*entities.User, error) {
	publicKey := j.getPublicKey()

	verifiedToken, err := jwt.Verify(jwt.EdDSA, publicKey, []byte(token))

	if err != nil {
		return nil, err
	}

	var claims jwt.Map
	_ = verifiedToken.Claims(&claims)

	userId, userIdExists := claims["user_id"]

	if !userIdExists {
		return nil, fmt.Errorf("user_id not found on token claims")
	}

	return entities.NewUser(userId.(string))
}

func NewJwtTokenVerifier(jwtConfig *config.JwtConfig) *JwtTokenVerifier {
	return &JwtTokenVerifier{
		config: jwtConfig,
	}
}
