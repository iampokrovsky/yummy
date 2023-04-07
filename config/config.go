package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	DB
}

type DB struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	DBName   string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	SSLMode  string `env:"DB_SSL"`
}

func (db *DB) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", db.Host, db.Port, db.DBName, db.User, db.Password, db.SSLMode)
}

func New() (Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return cfg, fmt.Errorf("config error: %w", err)
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return cfg, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
