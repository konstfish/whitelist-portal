package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RedisAddr        string `envconfig:"REDIS_ADDR" default:"localhost:6379"`
	AuthHeader       string `envconfig:"AUTH_HEADER" default:"X-Auth-Request-Email"`
	UserAddressLimit int    `envconfig:"USER_ADDRESS_LIMIT" default:"25"`
}

func Load() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
