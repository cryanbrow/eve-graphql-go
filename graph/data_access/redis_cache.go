package data_access

import (
	"strconv"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

func CheckRedisCache(key string) (bool, []byte) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == string("redis: nil") {
			return false, nil
		} else {
			log.Errorf("Redis encountered an error: %v", err)
		}
	}
	return true, []byte(val)
}

func AddToRedisCache(key string, value []byte, ttl int) {
	ttlString, err := time.ParseDuration((strconv.Itoa(ttl) + "ms"))
	if err != nil {
		log.Errorf("Failed to parse TTL: %v ", err)
		return
	}
	_, err = rdb.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == string("redis: nil") {
			rdb.Set(ctx, key, value, ttlString)
		} else {
			log.Errorf("Redis encountered an error: %v", err)
		}
	}
}

var (
	ctx context.Context = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
