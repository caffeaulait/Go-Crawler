package engine

import (
	"github.com/go-redis/redis"
	"log"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()
	if client == nil {
		t.Errorf("fail to create redis client")
	}

	//test connection
	pong, err := client.Ping().Result()
	if err != nil || pong != "PONG" {
		t.Errorf("failed to connect to redis")
	}

	err = client.Set("key", "value", 1*time.Second).Err()
	if err != nil {
		t.Errorf(" fail to set key")
	}

	ttl, err := client.TTL("key").Result()
	if err != nil {
		t.Errorf("fail to set ttl")
	} else {
		log.Printf("ttl is %v", ttl)
	}

	val, err := client.Get("missing_key").Result()
	if err == redis.Nil {
		log.Println("missing_key does not exist")
	} else if err == nil {
		t.Errorf("missing_key: %v", val)
	}
}
