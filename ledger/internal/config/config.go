package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPPort    string
	GRPCPort    string
	PostgresDSN string
	RedisAddr   string
	RedisPass   string
	RedisDB     int
}

func Load() Config {
	return Config{
		HTTPPort:    getEnv("HTTP_PORT", "8081"),
		GRPCPort:    getEnv("GRPC_PORT", "9091"),
		PostgresDSN: getEnv("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/ledger?sslmode=disable"),
		RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPass:   getEnv("REDIS_PASSWORD", ""),
		RedisDB:     getEnvInt("REDIS_DB", 0),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
