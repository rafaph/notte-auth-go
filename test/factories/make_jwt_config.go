package factories

import (
	"github.com/rafaph/notte-auth/config"
)

func MakeJwtConfig() *config.JwtConfig {
	conf, _ := config.NewConfig()

	return &conf.Jwt
}
