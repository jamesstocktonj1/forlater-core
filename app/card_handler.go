package app

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamesstocktonj1/forlater-core/internal/connect"
	"github.com/jamesstocktonj1/forlater-core/proto"
)

type CardHandler struct {
	conn   proto.CardServiceClient
	ctx    context.Context
	cancel context.CancelFunc
}

func NewCardHandler(config connect.ClientConfig) (*CardHandler, error) {
	var err error
	c := CardHandler{}

	c.ctx, c.cancel = context.WithTimeout(context.Background(), time.Millisecond*time.Duration(config.Timeout))
	conn, err := connect.NewClientConnection(config)
	if err != nil {
		return nil, err
	}

	c.conn = proto.NewCardServiceClient(conn)

	return &c, nil
}

func (u *CardHandler) HandleCreateCard(c *gin.Context) {
	card := CardData{}

	err := c.BindJSON(card)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to parse input json"})
		return
	}

	resp, err := u.conn.CreateCard(u.ctx, card.toProto())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal connection error"})
		return
	}

	switch resp.StatusCode {
	case proto.CardStatusCode_OK:
		c.IndentedJSON(http.StatusOK, resp)
		return
	case proto.CardStatusCode_ERROR:
	case proto.CardStatusCode_FORBIDDEN:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to create card"})
		return
	case proto.CardStatusCode_BAD_HASH:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid hash value"})
		return
	case proto.CardStatusCode_INTERNAL_ERROR:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.IndentedJSON(http.StatusTeapot, gin.H{"error": "how did you get here?"})
}

func (u *CardHandler) HandleSetCard(c *gin.Context) {
	card := CardData{}

	err := c.BindJSON(card)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to parse input json"})
		return
	}

	resp, err := u.conn.SetCard(u.ctx, card.toProto())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal connection error"})
		return
	}

	switch resp.StatusCode {
	case proto.CardStatusCode_OK:
		c.IndentedJSON(http.StatusOK, resp)
		return
	case proto.CardStatusCode_ERROR:
	case proto.CardStatusCode_FORBIDDEN:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to set card"})
		return
	case proto.CardStatusCode_BAD_HASH:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid hash value"})
		return
	case proto.CardStatusCode_INTERNAL_ERROR:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.IndentedJSON(http.StatusTeapot, gin.H{"error": "how did you get here?"})
}

func (u *CardHandler) HandleGetCard(c *gin.Context) {
	card := CardRequest{}

	err := c.BindJSON(card)
	singleQuery := c.Query("card")
	if singleQuery != "" {
		card.Cards = []string{singleQuery}
	} else if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "card query missing"})
		return
	}

	resp, err := u.conn.GetCard(u.ctx, card.toProto())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal connection error"})
		return
	}

	switch resp.StatusCode {
	case proto.CardStatusCode_OK:
		c.IndentedJSON(http.StatusOK, resp)
		return
	case proto.CardStatusCode_ERROR:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to get card"})
		return
	case proto.CardStatusCode_FORBIDDEN:
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "unauthorised"})
		return
	case proto.CardStatusCode_BAD_HASH:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid hash value"})
		return
	case proto.CardStatusCode_INTERNAL_ERROR:
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.IndentedJSON(http.StatusTeapot, gin.H{"error": "how did you get here?"})
}
