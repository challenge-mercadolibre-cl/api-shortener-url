package utils

import (
	"github.com/caarlos0/env/v6"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/infrastructure/log"
)

type ConfigServerHttp struct {
	Host string `env:"SERVER_HTTP_HOST"`
	Port string `env:"SERVER_HTTP_PORT"`
}

type ConfigRedis struct {
	Address  string `env:"REDIS_ADDRESS"`
	Password string `env:"REDIS_PASSWORD"`
	Database int    `env:"REDIS_DATABASE"`
	Username string `env:"REDIS_USERNAME"`
}
type Config struct {
	Server ConfigServerHttp
	Redis  ConfigRedis
}

func ReadConfig(opts ...env.Options) *Config {
	cfg := &Config{}
	if err := env.Parse(cfg, opts...); err != nil {
		log.WithError(err).Fatal("read config")
	}
	return cfg
}
