package authentication

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/jamesstocktonj1/forlater-core/internal/connect"
	"github.com/jamesstocktonj1/forlater-core/proto"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const (
	TestTimeout = 10
	TestAddr    = "0.0.0.0"
	TestToken   = "abcd"
)

var (
	mockRequest  *proto.TokenRequest
	mockResponse *proto.TokenResponse
	mockError    error
)

type mockUserService struct {
	proto.UserServiceClient
}

func (c *mockUserService) ValidateToken(ctx context.Context, in *proto.TokenRequest, opts ...grpc.CallOption) (*proto.TokenResponse, error) {
	mockRequest = in
	return mockResponse, mockError
}

func createMockAuthentication() (*Authentication, error) {
	a := Authentication{
		config: AuthenticationConfig{
			TokenTimeout: TestTimeout,
		},
	}

	mr, err := miniredis.Run()
	if err != nil {
		return nil, err
	}

	a.ctx = context.Background()
	a.cache = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	a.conn = &mockUserService{}
	mockRequest = nil
	mockResponse = nil
	mockError = nil

	return &a, nil
}

func createMockRequestServer(auth *Authentication) (*httptest.ResponseRecorder, *gin.Engine) {
	router := gin.New()
	router.SetTrustedProxies([]string{})
	router.Use(auth.Middleware())
	router.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	w := httptest.NewRecorder()

	return w, router
}

func createMockRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("authorization", "Bearer "+TestToken)
	return req
}

func TestNewAuthentication(t *testing.T) {
	config := AuthenticationConfig{
		TokenTimeout: TestTimeout,
	}
	userConfig := connect.ClientConfig{
		Addr:    TestAddr,
		Timeout: TestTimeout,
	}

	auth, err := NewAuthentication(config, userConfig)
	assert.NoError(t, err)
	assert.NotEmpty(t, auth)
}

func TestMiddleware(t *testing.T) {

	t.Run("valid_auth", func(t *testing.T) {
		auth, err := createMockAuthentication()
		assert.NoError(t, err)

		w, router := createMockRequestServer(auth)
		req := createMockRequest()

		auth.cache.Del(auth.ctx, "token:"+TestToken)

		mockError = nil
		mockResponse = &proto.TokenResponse{
			StatusCode: proto.StatusCode_STATUS_OK,
		}
		mockRequest = &proto.TokenRequest{}

		// test middleware
		router.ServeHTTP(w, req)

		assert.Equal(t, TestToken, mockRequest.Token)

		value, err := auth.cache.Get(auth.ctx, "token:"+TestToken).Result()
		assert.NoError(t, err)
		assert.Equal(t, "VALID", value)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("valid_auth_cache", func(t *testing.T) {
		auth, err := createMockAuthentication()
		assert.NoError(t, err)

		w, router := createMockRequestServer(auth)
		req := createMockRequest()

		auth.cache.Set(auth.ctx, "token:"+TestToken, TestToken, time.Minute)

		mockError = nil
		mockResponse = &proto.TokenResponse{
			StatusCode: proto.StatusCode_STATUS_OK,
		}
		mockRequest = &proto.TokenRequest{
			Token: "",
		}

		// test middleware
		router.ServeHTTP(w, req)

		assert.Equal(t, "", mockRequest.Token)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid_auth", func(t *testing.T) {
		auth, err := createMockAuthentication()
		assert.NoError(t, err)

		w, router := createMockRequestServer(auth)
		req := createMockRequest()

		auth.cache.Del(auth.ctx, "token:"+TestToken)

		mockError = nil
		mockResponse = &proto.TokenResponse{
			StatusCode: proto.StatusCode_STATUS_FORBIDDEN,
		}
		mockRequest = &proto.TokenRequest{
			Token: "",
		}

		// test middleware
		router.ServeHTTP(w, req)

		err = auth.cache.Get(auth.ctx, "token:"+TestToken).Err()
		assert.Equal(t, redis.Nil, err)
		assert.Equal(t, TestToken, mockRequest.Token)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("grpc_error", func(t *testing.T) {
		auth, err := createMockAuthentication()
		assert.NoError(t, err)

		w, router := createMockRequestServer(auth)
		req := createMockRequest()

		auth.cache.Del(auth.ctx, "token:"+TestToken)

		mockError = fmt.Errorf("Test Error")
		mockResponse = &proto.TokenResponse{
			StatusCode: proto.StatusCode_STATUS_FORBIDDEN,
		}
		mockRequest = &proto.TokenRequest{
			Token: "",
		}

		// test middleware
		router.ServeHTTP(w, req)

		err = auth.cache.Get(auth.ctx, "token:"+TestToken).Err()
		assert.Equal(t, redis.Nil, err)
		assert.Equal(t, TestToken, mockRequest.Token)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("missing_token", func(t *testing.T) {
		auth, err := createMockAuthentication()
		assert.NoError(t, err)

		w, router := createMockRequestServer(auth)
		req := createMockRequest()
		req.Header.Del("authorization")

		auth.cache.Del(auth.ctx, "token:"+TestToken)

		mockError = nil
		mockResponse = &proto.TokenResponse{
			StatusCode: proto.StatusCode_STATUS_FORBIDDEN,
		}
		mockRequest = &proto.TokenRequest{
			Token: "",
		}

		// test middleware
		router.ServeHTTP(w, req)

		err = auth.cache.Get(auth.ctx, "token:"+TestToken).Err()
		assert.Equal(t, redis.Nil, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
