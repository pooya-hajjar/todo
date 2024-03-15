package models

import (
	"log"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
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
	connectToRedisSearch()

	RedisDb = *rdb
}
