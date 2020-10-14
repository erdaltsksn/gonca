package database

import "github.com/go-redis/redis/v8"

// Redis returns a client to Redis Server.
func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0, // use default DB
	})

	return rdb
}
