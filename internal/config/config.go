package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTP struct {
	Port string `env:"HTTP_PORT" env-default:"8080"`
}

type DB struct {
	Name     string `env:"DB_NAME" env-default:"tokens"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
}

type Storage struct {
	UseMemory bool `env:"USE_MEMORY" env-default:"false"`
}

type Config struct {
	HTTP
	DB
	Storage
}

func New() (*Config, error) {

	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, errors.New("error reading env")
	}

	return &cfg, nil
}
