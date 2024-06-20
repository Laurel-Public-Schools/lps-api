package db

import (
	"context"
	"time"

	"github.com/laurel-public-schools/lps-api/env"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Redis *redis.ClusterClient
	ctx   context.Context
}

func NewRedisClient(ctx context.Context) (Client, error) {
	var (
		redisHost1 = env.GetConfig().REDIS_HOST1 + ":" + env.GetConfig().REDIS_PORT
		redisHost2 = env.GetConfig().REDIS_HOST2 + ":" + env.GetConfig().REDIS_PORT
		redisHost3 = env.GetConfig().REDIS_HOST3 + ":" + env.GetConfig().REDIS_PORT
	)
	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{redisHost1, redisHost2, redisHost3},
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return Client{}, err
	}

	return Client{
		Redis: rdb,
		ctx:   ctx,
	}, nil
}

func (c *Client) SetKV(key string, value string, expires time.Duration) error {
	err := c.Redis.Set(c.ctx, key, value, expires).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetKV(key string) string {
	val, err := c.Redis.Get(c.ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}
