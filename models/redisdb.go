package models

import (
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"sync"
)

var RedisDb redis.Client

func ConnectToRedis(wg *sync.WaitGroup) {
	defer wg.Done()

	redisUrl := os.Getenv("REDIS_URL")

	url, parseUrlErr := redis.ParseURL(redisUrl)
	if parseUrlErr != nil {
		log.Fatal("db connection error")
	}

	rdb := redis.NewClient(url)

	RedisDb = *rdb
}
