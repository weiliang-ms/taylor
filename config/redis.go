package config

import (
	"os"
)

type RedisInstance struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

const (
	//RedisDefaultHost = "127.0.0.1"
	RedisDefaultHost     = "192.168.40.146"
	RedisDefaultPort     = "6379"
	RedisDefaultPassword = "redis"
)

// redis信息
func RedisInstanceInfo() (instance RedisInstance) {
	instance = RedisInstance{
		Host:     RedisDefaultHost,
		Port:     RedisDefaultPort,
		Password: RedisDefaultPassword,
	}
	if host := os.Getenv("RedisHost"); host != "" {
		instance.Host = host
	}

	if port := os.Getenv("RedisPort"); port != "" {
		instance.Port = port
	}

	if password := os.Getenv("RedisPassword"); password != "" {
		instance.Password = password
	}
	return instance
}
