package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBUser string `env:"DB_USER,required"`
	DBPass string `env:"DB_PASS,required"`
	DBName string `env:"DB_NAME,required"`
	DBPort string `env:"DB_PORT,required"`
	DBHost string `env:"DB_HOST,required"`
}

func NewEnvConfig() *EnvConfig {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env: %e", err)
	}

	config := &EnvConfig{}
	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from .env: %e", err)
	}

	return config
}
