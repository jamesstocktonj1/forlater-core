package authentication

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamesstocktonj1/forlater-core/internal/connect"
	"github.com/jamesstocktonj1/forlater-core/internal/database"
	"github.com/jamesstocktonj1/forlater-core/proto"
	"github.com/redis/go-redis/v9"
)

type AuthenticationConfig struct {
	database.CacheConfig
	TokenTimeout int `json:"timeout"`
}

type Authentication struct {
	config AuthenticationConfig
	cache  *redis.Client
	conn   proto.UserServiceClient
	ctx    context.Context
	cancel context.CancelFunc
}

func NewAuthentication(config AuthenticationConfig, userConfig connect.ClientConfig) (*Authentication, error) {
	a := Authentication{
		config: config,
	}

	a.ctx, a.cancel = context.WithTimeout(context.Background(), time.Millisecond*time.Duration(userConfig.Timeout))
	conn, err := connect.NewClientConnection(userConfig)
	if err != nil {
		return nil, err
	}

	a.conn = proto.NewUserServiceClient(conn)
	a.cache = database.NewCache(config.CacheConfig)

	return &a, nil
}

func (a *Authentication) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "authorization header not found"})
			c.Abort()
			return
		}

		rawToken, _ := strings.CutPrefix(token, "Bearer ")

		err := a.cache.Get(a.ctx, "token:"+rawToken).Err()
		if err == nil {
			c.Next()
			return
		}

		resp, err := a.conn.ValidateToken(a.ctx, &proto.TokenRequest{Token: rawToken})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to validate token"})
			c.Abort()
			return
		}

		if resp.StatusCode != proto.StatusCode_STATUS_OK {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired auth token"})
			c.Abort()
			return
		}

		a.cacheUserToken(rawToken)
		c.Next()
	}
}

func (a *Authentication) cacheUserToken(rawToken string) {
	err := a.cache.SetEx(a.ctx, "token:"+rawToken, "VALID", time.Second*time.Duration(a.config.TokenTimeout)).Err()
	if err != nil {
		log.Println("Authentication Error: failed to cache auth token")
	}
}
