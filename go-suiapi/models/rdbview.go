package models

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func RDBhasViewed(redisKey string) bool {
	val, err := rdb.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		log.Printf("redis err: %s", err)
	}
	return val == "1"
}

func RDBmarkAsViewed(redisKey string, expiration time.Duration) {
	err := rdb.Set(ctx, redisKey, "1", expiration).Err()
	if err != nil {
		log.Printf("redis err: %s", err)
	}
}
