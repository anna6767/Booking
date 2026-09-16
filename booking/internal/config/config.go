package config

import "os"

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

func LoadConfig() Config {
	var cfg Config

	cfg.DBHost = os.Getenv("DB_HOST")
	if cfg.DBHost == "" {
		cfg.DBHost = "localhost"
	}

	cfg.DBPort = os.Getenv("DB_PORT")
	if cfg.DBPort == "" {
		cfg.DBPort = "5432"
	}

	cfg.DBUser = os.Getenv("DB_USER")
	if cfg.DBUser == "" {
		cfg.DBUser = "bookinguser"
	}

	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	if cfg.DBPassword == "" {
		cfg.DBPassword = "secretpass"
	}

	cfg.DBName = os.Getenv("DB_NAME")
	if cfg.DBName == "" {
		cfg.DBName = "bookingdb"
	}

	cfg.AppPort = os.Getenv("APP_PORT")
	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}

	return cfg
}
