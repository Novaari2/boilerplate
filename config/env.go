package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/rs/zerolog/log"
)

func Load() *Config {
	var c Config
	if err := env.Parse(&c); err != nil {
		log.Fatal().Msgf("failed to load config: %v", err.Error())
	}

	return &c
}

type Config struct {
	Database Database
}

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}
