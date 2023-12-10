package ratelimit

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamesstocktonj1/forlater-core/internal/database"
	"github.com/redis/go-redis/v9"
)

type RateLimitConfig struct {
	Timeout int   `json:"timeout"`
	Count   int64 `json:"count"`
}

type RateLimit struct {
	cache  *redis.Client
	ctx    context.Context
	config RateLimitConfig
}

func NewRateLimit(config RateLimitConfig, cacheConfig database.CacheConfig) RateLimit {
	r := RateLimit{
		config: config,
	}

	r.ctx = context.Background()
	r.cache = database.NewCache(cacheConfig)

	log.Printf("Rate Limit Created with Timeout: %d and Count: %d\n", config.Timeout, config.Count)

	return r
}

func (r *RateLimit) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sourceIP := c.ClientIP()

		ok, err := r.incrementCounter(sourceIP)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "proxy error"})
			c.Abort()
		} else if !ok {
			c.IndentedJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func (r *RateLimit) incrementCounter(addr string) (bool, error) {
	limitKey := "ratelimit:" + addr

	log.Printf("Rate Limit: %s\n", limitKey)

	count, err := r.cache.Incr(r.ctx, limitKey).Result()
	if err != nil {
		return false, err
	}

	expireTime := time.Second * time.Duration(r.config.Timeout)
	err = r.cache.Expire(r.ctx, limitKey, expireTime).Err()
	if err != nil {
		return false, err
	}

	return count < r.config.Count, nil
}
