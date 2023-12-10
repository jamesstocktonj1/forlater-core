package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamesstocktonj1/forlater-core/internal/database"
	"github.com/jamesstocktonj1/forlater-core/middleware/ratelimit"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	gin    *gin.Engine
	ctx    context.Context
	cache  *redis.Client
	config ServerConfig
}

func NewServer(config ServerConfig) Server {
	s := Server{
		config: config,
	}

	s.gin = gin.Default()

	rateLimiter := ratelimit.NewRateLimit(config.Ratelimiter, config.Redis)
	s.gin.Use(rateLimiter.Middleware())
	s.gin.GET("/ping", s.Ping)

	s.ctx = context.Background()
	s.cache = database.NewCache(config.Redis)

	return s
}

func (s *Server) Run() error {
	err := s.cache.Ping(s.ctx).Err()
	if err != nil {
		return fmt.Errorf("Unable to connect to Redis Cache: %s", err.Error())
	}

	log.Printf("Hosting on Address: %s\n", s.config.HttpsAddr)
	return s.gin.Run(s.config.HttpsAddr)
}

func (s *Server) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Pong"})
}
