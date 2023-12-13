package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseHost     string `envconfig:"DB_HOST" required:"true"`
	DatabasePort     string `envconfig:"DB_PORT" default:"3306"`
	DatabaseUser     string `envconfig:"DB_USER" required:"true"`
	DatabasePassword string `envconfig:"DB_PASSWORD" required:"true"`
	DatabaseName     string `envconfig:"DB_NAME" required:"true"`
	MaxOpenConns     int    `envconfig:"MAX_OPEN_CONNS" default:"5"`
	MaxIdleConns     int    `envconfig:"MAX_IDLE_CONNS" default:"5"`
}

func ParseEnv() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("config env: unable to pars env vars: %w", err)
	}
	return &cfg, nil
}
