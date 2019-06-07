package config

import (
	_ "go-phonebooks/utils/env"
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

var config *Config

func init() {
	config = &Config{
		DB: &DBConfig{
			Dialect:  os.Getenv("db_type"),
			Name:     os.Getenv("db_name"),
			Port:     os.Getenv("db_port"),
			Username: os.Getenv("db_user"),
			Password: os.Getenv("db_password"),
			Host:     os.Getenv("db_host"),
			Charset:  "utf8",
			Loc:      "Local",
		},
	}
}

func GetConfig() *Config {
	return config
}
