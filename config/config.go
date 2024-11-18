package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	ENV         string      `env:"ENV"     envDefault:"dev"`
	PORT        string      `env:"PORT"    envDefault:"8080"`
	MySQLConfig MySQLConfig `envPrefix:"MYSQL_"`
}

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST"     envDefault:"localhost"`
	Port     string `env:"MYSQL_PORT"    envDefault:"3306"`
	User     string `env:"MYSQL_USER"     envDefault:"root"`
	Password string `env:"MYSQL_PASSWORD" envDefault:""`
	Database string `env:"MYSQL_DATABASE" envDefault:""`
}

func NewConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
