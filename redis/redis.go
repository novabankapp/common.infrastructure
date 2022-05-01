package redis

import (
	"github.com/go-redis/redis/v8"
)

func NewUniversalRedisClient(cfg *Config) redis.UniversalClient {

	universalClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:           []string{cfg.Addr},
		Password:        cfg.Password, // no password set
		DB:              cfg.DB,       // use default DB
		MaxRetries:      maxRetries,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		DialTimeout:     dialTimeout,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		PoolSize:        cfg.PoolSize,
		MinIdleConns:    minIdleConns,
		PoolTimeout:     poolTimeout,
		IdleTimeout:     idleTimeout,
	})

	return universalClient
}
