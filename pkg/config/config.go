package config

import (
	"drones/pkg/env"
	"sync"
)

var (
	once         sync.Once
	confInstance *Config
)

type Config struct {
	DB struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
		SSlMode  string
	}
	Server struct {
		Host string
		Port string
	}
}

func Construct() {
	once.Do(func() {
		confInstance = &Config{
			DB: struct {
				User     string
				Password string
				Host     string
				Port     string
				Name     string
				SSlMode  string
			}{
				User:     env.String("DB_USER", "postgres"),
				Password: env.String("DB_PASSWORD", "pass"),
				Host:     env.String("DB_HOST", "localhost"),
				Port:     env.String("DB_PORT", "5432"),
				Name:     env.String("DB_NAME", "drones"),
				SSlMode:  env.String("DB_SSL", "disable"),
			},
			Server: struct {
				Host string
				Port string
			}{
				Host: env.String("SERVER_HOST", "localhost"),
				Port: env.String("SERVER_PORT", "6504"),
			},
		}
	})
}

func GetConfig() *Config {
	if confInstance == nil {
		Construct()
	}

	return confInstance
}
