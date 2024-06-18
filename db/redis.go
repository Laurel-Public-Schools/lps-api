package db

import (
	"github.com/laurel-public-schools/lps-api/env"
	"github.com/redis/go-redis/v9"
)

func Redis() *redis.ClusterClient {
	var (
		redisHost1 = env.GetConfig().REDIS_HOST1 + ":" + env.GetConfig().REDIS_PORT
		redisHost2 = env.GetConfig().REDIS_HOST2 + ":" + env.GetConfig().REDIS_PORT
		redisHost3 = env.GetConfig().REDIS_HOST3 + ":" + env.GetConfig().REDIS_PORT
	)
	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{redisHost1, redisHost2, redisHost3},
	})
	return rdb
}
