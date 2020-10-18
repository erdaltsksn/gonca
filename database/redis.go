package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// Redis returns a client to Redis Server.
func Redis() *redis.Client {
	options := &redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0, // Default DB
	}

	if viper.IsSet("redis.addr") {
		options.Addr = viper.GetString("redis.addr")
	}
	if viper.IsSet("redis.pass") {
		options.Password = viper.GetString("redis.pass")
	}
	if viper.IsSet("redis.db") {
		options.DB = viper.GetInt("redis.db")
	}

	rdb := redis.NewClient(options)

	return rdb
}
