package app

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamesstocktonj1/forlater-core/internal/connect"
	"github.com/jamesstocktonj1/forlater-core/proto"
)

type UserHandler struct {
	conn   proto.UserServiceClient
	ctx    context.Context
	cancel context.CancelFunc
	config connect.ClientConfig
}

func NewUserHandler(config connect.ClientConfig) (*UserHandler, error) {
	var err error
	u := UserHandler{}

	u.ctx, u.cancel = context.WithTimeout(context.Background(), time.Millisecond*time.Duration(config.Timeout))
	conn, err := connect.NewClientConnection(config)
	if err != nil {
		return nil, err
	}

	u.conn = proto.NewUserServiceClient(conn)
	u.config = config

	return &u, nil
}

func (u *UserHandler) HandleCreateUser(c *gin.Context) {
	user := UserData{}

	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to parse input json"})
		return
	}

	u.ctx, u.cancel = context.WithTimeout(u.ctx, time.Millisecond*time.Duration(u.config.Timeout))
	resp, err := u.conn.CreateUser(u.ctx, user.toProto())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal connection error"})
		return
	}

	switch resp.StatusCode {
	case proto.StatusCode_STATUS_OK:
		c.IndentedJSON(http.StatusOK, toUser(resp))
		return
	case proto.StatusCode_STATUS_ERROR:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to create user"})
		return
	case proto.StatusCode_STATUS_INTERNAL_ERROR:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.IndentedJSON(http.StatusTeapot, gin.H{"error": "how did you get here?"})
}

func (u *UserHandler) HandleSetUser(c *gin.Context) {
	user := UserData{}

	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to parse input json"})
		return
	}

	u.ctx, u.cancel = context.WithTimeout(u.ctx, time.Millisecond*time.Duration(u.config.Timeout))
	resp, err := u.conn.SetUser(u.ctx, user.toProto())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal connection error"})
		return
	}

	userResponse := toUser(resp)

	switch resp.StatusCode {
	case proto.StatusCode_STATUS_OK:
		c.IndentedJSON(http.StatusOK, userResponse)
		return
	case proto.StatusCode_STATUS_ERROR:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to set user"})
		return
	case proto.StatusCode_STATUS_INTERNAL_ERROR:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.IndentedJSON(http.StatusTeapot, gin.H{"error": "how did you get here?"})
}

func (u *UserHandler) HandleGetUser(c *gin.Context) {
	user := UserData{}

	user.Username = c.Query("username")
	if user.Username == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "username query missing"})
		return
	}

	u.ctx, u.cancel = context.WithTimeout(u.ctx, time.Millisecond*time.Duration(u.config.Timeout))
	resp, err := u.conn.GetUser(u.ctx, user.toProto())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal connection error"})
		return
	}

	switch resp.StatusCode {
	case proto.StatusCode_STATUS_OK:
		c.IndentedJSON(http.StatusOK, toUser(resp))
		return
	case proto.StatusCode_STATUS_ERROR:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to set user"})
		return
	case proto.StatusCode_STATUS_INTERNAL_ERROR:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.IndentedJSON(http.StatusTeapot, gin.H{"error": "how did you get here?"})
}

func (u *UserHandler) HandleLoginUser(c *gin.Context) {
	user := UserData{}

	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to parse input json"})
		return
	}

	u.ctx, u.cancel = context.WithTimeout(u.ctx, time.Millisecond*time.Duration(u.config.Timeout))
	resp, err := u.conn.LoginUser(u.ctx, user.toProto())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal connection error"})
		return
	}

	tokenResponse := toToken(resp)
	tokenResponse.Username = user.Username

	switch resp.StatusCode {
	case proto.StatusCode_STATUS_OK:
		c.IndentedJSON(http.StatusOK, tokenResponse)
		return
	case proto.StatusCode_STATUS_ERROR:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to login"})
		return
	case proto.StatusCode_STATUS_INTERNAL_ERROR:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.IndentedJSON(http.StatusTeapot, gin.H{"error": "how did you get here?"})
}
