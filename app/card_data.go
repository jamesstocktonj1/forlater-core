package app

import "github.com/jamesstocktonj1/forlater-core/proto"

type CardData struct {
	CardID    string `json:"card_id"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
	Hash      string `json:"hash"`
}

type CardRequest struct {
	Username string   `json:"username"`
	Cards    []string `json:"cards"`
}

func (c *CardData) toProto() *proto.Card {
	return &proto.Card{
		CardId:    c.CardID,
		Username:  c.Username,
		Content:   c.Content,
		Timestamp: c.Timestamp,
		Hash:      c.Hash,
	}
}

func (c *CardRequest) toProto() *proto.CardRequest {
	return &proto.CardRequest{
		Username: c.Username,
		CardId:   c.Cards,
	}
}
