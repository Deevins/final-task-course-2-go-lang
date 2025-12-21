package config

import "os"

type Config struct {
	HTTPPort    string
	GRPCPort    string
	PostgresDSN string
	JWT         JWTConfig
}

func Load() (Config, error) {
	jwtConfig, err := LoadJWTConfig()
	if err != nil {
		return Config{}, err
	}

	return Config{
		HTTPPort:    getEnv("HTTP_PORT", "8082"),
		GRPCPort:    getEnv("GRPC_PORT", "9092"),
		PostgresDSN: getEnv("AUTH_POSTGRES_DSN", ""),
		JWT:         jwtConfig,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
