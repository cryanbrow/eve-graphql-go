package caching

import (
	"context"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/go-redis/redis/v8"
)

//Caching interfaces for implementing caching, for retrieving records, and checking for existence, and adding to a cache with a TTL
type Caching interface {
	CheckCache(ctx context.Context, key string) (bool, []byte)
	AddToCache(ctx context.Context, key string, value []byte, ttl int64)
}

var (
	//Cache implementation. Injected object will have CheckCache and AddToCache available.
	Cache Caching
)

//ConfigureCaching is used to setup caching, if the configuration caching impl is redis, then redis will be used.
//If any other value is provided including empty it will be an in memory cache.
//Default is in memory. If using redis, implement configs for URL, Port, Username, and Password.
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
