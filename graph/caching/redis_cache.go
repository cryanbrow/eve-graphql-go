package caching

import (
	"strconv"
	"time"

	"context"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type Client struct {
}

const tracer_name = "github.com/cryanbrow/eve-graphql-go/graph/caching"

func (c *Client) CheckRedisCache(key string, ctx context.Context) (bool, []byte) {
	_, span := otel.Tracer(tracer_name).Start(ctx, "CheckRedisCache")
	defer span.End()
	log.Debugf("Checking Redis Cache for key: %s", key)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == string("redis: nil") {
			return false, nil
		} else {
			log.Errorf("Redis encountered an error: %v", err)
		}
	}
	span.SetAttributes(attribute.String("key", key))
	return true, []byte(val)
}

func (c *Client) AddToRedisCache(key string, value []byte, ttl int64, ctx context.Context) {
	_, span := otel.Tracer(tracer_name).Start(ctx, "AddToRedisCache")
	defer span.End()
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
			statusText, _ := status.Result()
			log.Debugf("status text: %s", statusText)
		} else {
			log.Errorf("Redis encountered an error: %v", err)
		}
	}
	span.SetAttributes(attribute.String("key", key), attribute.Int64("ttl", ttl))
}

func ConfigureRedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     configuration.AppConfig.Redis.Url + ":" + configuration.AppConfig.Redis.Port,
		Username: configuration.AppConfig.Redis.User,
		Password: configuration.AppConfig.Redis.Password,
		DB:       0, // use default DB
	})
}

var (
	rdb *redis.Client
)
