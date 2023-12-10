package app

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jamesstocktonj1/forlater-core/internal/database"
	"github.com/jamesstocktonj1/forlater-core/middleware/ratelimit"
)

type ServerConfig struct {
	HttpsAddr   string                    `json:"https_addr"`
	Redis       database.CacheConfig      `json:"redis"`
	Ratelimiter ratelimit.RateLimitConfig `json:"rate_limit"`
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
