package database

import (
	"github.com/redis/go-redis/v9"
)

type CacheConfig struct {
	Addr string `json:"redis_addr"`
}

func NewCache(config CacheConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})
}
