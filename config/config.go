package config

import (
	"os"
)

type DBConfig struct {
	Dialect  string
	Name     string
	Port     string
	Username string
	Password string
	Host     string
	Charset  string
	Loc      string
}

type Config struct {
	DB *DBConfig
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  os.Getenv("DB_DIALECT"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Charset:  os.Getenv("DB_CHARSET"),
			Loc:      "Local",
		},
	}
}
