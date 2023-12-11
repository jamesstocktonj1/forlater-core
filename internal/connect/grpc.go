package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConfig struct {
	Addr    string `json:"addr"`
	Timeout int    `json:"timeout"`
}

func NewClientConnection(config ClientConfig) (*grpc.ClientConn, error) {
	return grpc.Dial(config.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
