package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context, addr, password string, db int) (*redis.Client, error) {
	if addr == "" {
		return nil, fmt.Errorf("redis addr is required")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		if closeErr := client.Close(); closeErr != nil {
			return nil, fmt.Errorf("ping redis: %w (close error: %v)", err, closeErr)
		}
		return nil, fmt.Errorf("ping redis: %w", err)
	}
	return client, nil
}
