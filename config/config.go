package config

import (
	"github.com/rafaph/notte-auth/lib/validator"
	"os"
	"strconv"
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

	env, ok := os.LookupEnv(key)

	if ok {
		transformedEnv, err := transform(env)
		if err == nil {
			*value = transformedEnv
		}
	}
}

type JwtConfig struct {
	ExpirationInMinutes int    `validate:"required,gt=0"`
	Secret              string `validate:"required"`
}

func newJwtConfig() JwtConfig {
	conf := JwtConfig{}

	setValueFromEnv(&conf.ExpirationInMinutes, "JWT_EXPIRATION_IN_MINUTES", strconv.Atoi)
	setValueFromEnv(&conf.Secret, "JWT_SECRET")

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
