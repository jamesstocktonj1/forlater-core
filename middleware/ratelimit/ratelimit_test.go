package ratelimit

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

const (
	TestCount   = 10
	TestTimeout = 10
	TestAddr    = "0.0.0.0"
)

func createMockRateLimit() (*RateLimit, error) {
	r := RateLimit{
		config: RateLimitConfig{
			Count:   TestCount,
			Timeout: TestTimeout,
		},
	}

	mr, err := miniredis.Run()
	if err != nil {
		return nil, err
	}

	r.ctx = context.Background()
	r.cache = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return &r, nil
}

func createMockRequestServer(limit *RateLimit) (*httptest.ResponseRecorder, *gin.Engine) {
	router := gin.New()
	router.SetTrustedProxies([]string{})
	router.Use(limit.Middleware())
	router.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	w := httptest.NewRecorder()

	return w, router
}

func createMockRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.RemoteAddr = TestAddr + ":0000"
	return req
}

func TestMiddleware(t *testing.T) {

	t.Run("new_addr", func(t *testing.T) {
		r, err := createMockRateLimit()
		assert.NoError(t, err)

		w, router := createMockRequestServer(r)
		req := createMockRequest()

		// delete rate limit key
		r.cache.Del(r.ctx, "ratelimit:"+TestAddr)

		router.ServeHTTP(w, req)

		value, err := r.cache.Get(r.ctx, "ratelimit:"+TestAddr).Int()
		assert.NoError(t, err)
		assert.Equal(t, 1, value)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("increment_existing_addr", func(t *testing.T) {
		r, err := createMockRateLimit()
		assert.NoError(t, err)

		w, router := createMockRequestServer(r)
		req := createMockRequest()

		r.cache.Set(r.ctx, "ratelimit:"+TestAddr, 1, time.Minute)

		router.ServeHTTP(w, req)

		value, err := r.cache.Get(r.ctx, "ratelimit:"+TestAddr).Int()
		assert.NoError(t, err)
		assert.Equal(t, 2, value)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("increment_before_limit", func(t *testing.T) {
		r, err := createMockRateLimit()
		assert.NoError(t, err)

		w, router := createMockRequestServer(r)
		req := createMockRequest()

		r.cache.Set(r.ctx, "ratelimit:"+TestAddr, TestCount-2, time.Minute)

		router.ServeHTTP(w, req)

		value, err := r.cache.Get(r.ctx, "ratelimit:"+TestAddr).Int()
		assert.NoError(t, err)
		assert.Equal(t, TestCount-1, value)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("increment_before_limit", func(t *testing.T) {
		r, err := createMockRateLimit()
		assert.NoError(t, err)

		w, router := createMockRequestServer(r)
		req := createMockRequest()

		r.cache.Set(r.ctx, "ratelimit:"+TestAddr, TestCount-1, time.Minute)

		router.ServeHTTP(w, req)

		value, err := r.cache.Get(r.ctx, "ratelimit:"+TestAddr).Int()
		assert.NoError(t, err)
		assert.Equal(t, TestCount, value)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})
}
