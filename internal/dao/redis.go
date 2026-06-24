package dao

import (
	"context"
	"fmt"

	"gomind/internal/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis(ctx context.Context, cfg config.RedisConfig, addr string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.Database,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		if closeErr := client.Close(); closeErr != nil {
			return nil, fmt.Errorf("ping redis: %w; close redis client: %v", err, closeErr)
		}
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	return client, nil
}
