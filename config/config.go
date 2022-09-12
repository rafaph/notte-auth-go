package config

import (
	"github.com/rafaph/notte-auth/lib/validator"
	"os"
	"strconv"
	"strings"
)

func setValueFromEnv[T any](value *T, key string, transforms ...func(value string) (T, error)) {
	var transform func(value string) (T, error)

	if len(transforms) > 0 {
		transform = transforms[0]
	} else {
		transform = func(value string) (T, error) {
			var transformedEnv interface{} = value
			return transformedEnv.(T), nil
		}
	}

	if env, ok := os.LookupEnv(key); ok {
		if transformedEnv, err := transform(env); err == nil {
			*value = transformedEnv
		}
	}
}

func transformJwtKey(value string) (string, error) {
	return strings.ReplaceAll(value, "\\n", "\n"), nil
}

type JwtConfig struct {
	ExpirationInMinutes int    `validate:"required,gt=0"`
	PrivateKey          string `validate:"required"`
	PublicKey           string `validate:"required"`
}

func newJwtConfig() JwtConfig {
	conf := JwtConfig{}

	setValueFromEnv(&conf.ExpirationInMinutes, "JWT_EXPIRATION_IN_MINUTES", strconv.Atoi)
	setValueFromEnv(&conf.PrivateKey, "JWT_PRIVATE_KEY", transformJwtKey)
	setValueFromEnv(&conf.PublicKey, "JWT_PUBLIC_KEY", transformJwtKey)

	return conf
}

type Config struct {
	Jwt JwtConfig `validate:"required,dive,required"`
}

func NewConfig() (*Config, error) {
	config := &Config{
		Jwt: newJwtConfig(),
	}

	err := validator.Validate(config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
