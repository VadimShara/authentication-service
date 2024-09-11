package jwt

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	Secret                string        `env:"JWT_SECRET"`
	AccessTokenExpiration time.Duration `env:"JWT_ACCESS_TOKEN_EXPIRATION" envDefault:"15m"`
	RefreshTokenExpiration time.Duration `env:"JWT_REFRESH_TOKEN_EXPIRATION" envDefault:"7d"`
}

func ParseConfig() (*Config, error) {
	var c Config
	if err := env.Parse(&c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &c, nil
}