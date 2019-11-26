package dao

import "github.com/go-redis/redis/v7"

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "49.235.67.172:6379",
		Password: "redis-bazinga@1qaz2wsx", // no password set
		DB:       0,                        // use default DB
	})
}
