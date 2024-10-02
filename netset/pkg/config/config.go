package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RedisAddr string `envconfig:"REDIS_ADDR" default:"localhost:6379"`
}

func Load() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
