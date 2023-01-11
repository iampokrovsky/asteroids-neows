package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP  `yaml:"http"`
	NeoWs `yaml:"neows"`
	DB    `yaml:"db"`
}

type HTTP struct {
	Port string `yaml:"port"`
}

type NeoWs struct {
	URL    string `yaml:"url"`
	ApiKey string `yaml:"api_key"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
