package redis

import (
	"context"
	"time"

	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) port.CacheRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *redisRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
