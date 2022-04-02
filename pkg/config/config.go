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
	Server struct {
		Host string
		Port string
	}
}

func Construct() {
	once.Do(func() {
		confInstance = &Config{
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
