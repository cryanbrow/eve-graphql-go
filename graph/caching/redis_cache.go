package caching

import (
	"strconv"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// RedisClient holds values related to connecting and using redis for caching.
type RedisClient struct {
}

const tracerName = "github.com/cryanbrow/eve-graphql-go/graph/caching"

// CheckCache takes in a context for tracing and a key for searching redis. Returns whether it found the record and the []byte of the record.
func (c *RedisClient) CheckCache(ctx context.Context, key string) (bool, []byte) {
	_, span := otel.Tracer(tracerName).Start(ctx, "CheckRedisCache")
	defer span.End()
	log.Debugf("Checking Redis Cache for key: %s", key)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == string("redis: nil") {
			return false, nil
		}
		log.Errorf("Redis encountered an error: %v", err)
	}
	span.SetAttributes(attribute.String("key", key))
	return true, []byte(val)
}

// AddToCache takes in a context for tracing, a key of the record to be stored, a []byte to be stored, and a ttl in milliseconds for
// how long the record should live.
func (c *RedisClient) AddToCache(ctx context.Context, key string, value []byte, ttl int64) {
	_, span := otel.Tracer(tracerName).Start(ctx, "AddToRedisCache")
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

var (
	rdb *redis.Client
)
