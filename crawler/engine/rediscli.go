package engine

import (
	"github.com/go-redis/redis"
)

func CreateRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func IsDuplicate(client *redis.Client, url string) bool {
	_, err := client.Get(url).Result()
	if err == redis.Nil {
		client.Set(url, "ok", 0)
		return false
	} else if err == nil {
		return true
	} else {
		panic(err)
	}
}
