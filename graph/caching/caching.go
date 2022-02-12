package caching

import (
	"context"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/go-redis/redis/v8"
)

type Caching interface {
	CheckCache(key string, ctx context.Context) (bool, []byte)
	AddToCache(key string, value []byte, ttl int64, ctx context.Context)
}

var (
	Cache Caching
)

func ConfigureCaching() {
	if configuration.AppConfig.Caching.Impl == "redis" {
		rdb = redis.NewClient(&redis.Options{
			Addr:     configuration.AppConfig.Redis.Url + ":" + configuration.AppConfig.Redis.Port,
			Username: configuration.AppConfig.Redis.User,
			Password: configuration.AppConfig.Redis.Password,
			DB:       0, // use default DB
		})
		Cache = &RedisClient{}
	} else {
		Cache = &MemoryClient{}
	}
}
