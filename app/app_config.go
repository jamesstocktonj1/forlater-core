package app

import (
	"encoding/json"
	"io"
	"os"
)

type ServerConfig struct {
	HttpsAddr string `json:"https_addr"`
	RedisAddr string `json:"redis_addr"`
}

func LoadConfig(filename string) (ServerConfig, error) {
	var config ServerConfig

	file, err := os.Open(filename)
	if err != nil {
		return ServerConfig{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return ServerConfig{}, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return ServerConfig{}, err
	}

	return config, nil
}
