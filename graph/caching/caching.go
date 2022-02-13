package caching

import (
	"context"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/go-redis/redis/v8"
)

type Caching interface {
	CheckCache(ctx context.Context, key string) (bool, []byte)
	AddToCache(ctx context.Context, key string, value []byte, ttl int64)
}

var (
	Cache Caching
)

func ConfigureCaching() {
	if configuration.AppConfig.Caching.Impl == "redis" {
		rdb = redis.NewClient(&redis.Options{
			Addr:     configuration.AppConfig.Redis.URL + ":" + configuration.AppConfig.Redis.Port,
			Username: configuration.AppConfig.Redis.User,
			Password: configuration.AppConfig.Redis.Password,
			DB:       0, // use default DB
		})
		Cache = &RedisClient{}
	} else {
		Cache = &MemoryClient{}
	}
}
