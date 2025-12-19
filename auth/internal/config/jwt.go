package config

import (
	"errors"
	"os"
	"time"
)

const DefaultJWTExpiry = 24 * time.Hour

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

func LoadJWTConfig() (JWTConfig, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return JWTConfig{}, errors.New("JWT_SECRET is required")
	}

	expiryValue := os.Getenv("JWT_EXPIRY")
	if expiryValue == "" {
		return JWTConfig{Secret: secret, Expiry: DefaultJWTExpiry}, nil
	}

	expiry, err := time.ParseDuration(expiryValue)
	if err != nil {
		return JWTConfig{}, err
	}
	if expiry <= 0 {
		return JWTConfig{}, errors.New("JWT_EXPIRY must be positive")
	}

	return JWTConfig{Secret: secret, Expiry: expiry}, nil
}
