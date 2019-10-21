package utils

import (
	"github.com/go-redis/redis/v7"
	"taylor/config"
)

//TODO:后续待添加连接池
func RedisClient() *redis.Client {
	instance := config.RedisInstanceInfo()
	client := redis.NewClient(&redis.Options{
		Addr:     instance.Host + ":" + instance.Port,
		Password: instance.Password, // no password set
		DB:       0,                 // use default DB
	})
	pong, err := client.Ping().Result()
	Logger.Info(pong, err)

	return client
}
