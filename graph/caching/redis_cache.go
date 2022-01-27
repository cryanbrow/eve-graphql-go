package caching

import (
	"strconv"
	"time"

	"context"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

func CheckRedisCache(key string) (bool, []byte) {
	log.Debugf("Checking Redis Cache for key: %s", key)
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

func AddToRedisCache(key string, value []byte, ttl int64) {
	log.Debugf("Adding to Redis Cache: %s", key)
	ttlString, err := time.ParseDuration((strconv.FormatInt(ttl, 10) + "ms"))
	if err != nil {
		log.Errorf("Failed to parse TTL: %v ", err)
		return
	}
	_, err = rdb.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == string("redis: nil") {
			status := rdb.Set(ctx, key, value, ttlString)
			statusText, err := status.Result()
			log.Errorf("status text: %s Error: %v", statusText, err)
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
		Addr:     configuration.AppConfig.Redis.Url + ":" + configuration.AppConfig.Redis.Port,
		Username: configuration.AppConfig.Redis.User,
		Password: configuration.AppConfig.Redis.Password,
		DB:       0, // use default DB
	})
}
