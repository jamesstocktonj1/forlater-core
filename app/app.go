package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamesstocktonj1/forlater-core/internal/config"
	"github.com/jamesstocktonj1/forlater-core/internal/database"
	"github.com/jamesstocktonj1/forlater-core/middleware/authentication"
	"github.com/jamesstocktonj1/forlater-core/middleware/ratelimit"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	gin    *gin.Engine
	ctx    context.Context
	cache  *redis.Client
	config *config.ServerConfig
	consul *ConsulClient
	secure *gin.RouterGroup
}

func NewServer() Server {
	var err error
	s := Server{}

	// Consul Config Init
	s.consul, err = NewConsulClient()
	if err != nil {
		log.Fatal(err)
	}

	err = s.consul.Register()
	if err != nil {
		log.Fatal(err)
	}
	defer s.consul.Deregister()

	s.config, err = s.consul.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Gin API Init
	s.gin = gin.Default()

	rateLimiter := ratelimit.NewRateLimit(s.config)
	s.gin.Use(rateLimiter.Middleware())
	s.gin.GET("/ping", s.Ping)

	authentication, err := authentication.NewAuthentication(s.config.Authentication, s.config.UserService)
	if err != nil {
		log.Fatal(err)
	}
	s.secure = s.gin.Group("/")
	s.secure.Use(authentication.Middleware())

	userService, err := NewUserHandler(s.config.UserService)
	if err != nil {
		log.Fatal(err)
	}
	s.gin.POST("/user", userService.HandleCreateUser)
	s.gin.POST("/login", userService.HandleLoginUser)
	s.secure.PUT("/user", userService.HandleSetUser)
	s.secure.GET("/user", userService.HandleGetUser)

	cardService, err := NewCardHandler(s.config.CardService)
	if err != nil {
		log.Fatal(err)
	}
	s.secure.POST("/card", cardService.HandleCreateCard)
	s.secure.PUT("/card", cardService.HandleSetCard)
	s.secure.GET("/card", cardService.HandleGetCard)

	s.ctx = context.Background()
	s.cache = database.NewCache(s.config.Cache)

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
