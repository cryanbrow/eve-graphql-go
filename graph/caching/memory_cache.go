package caching

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// MemoryClient for implementing in memory cache.
type MemoryClient struct {
}

type cacheRecord struct {
	value  []byte
	expiry int64
}

var memoryCache sync.Map

// CheckCache method for checking if the cache has a key matching the string provided. The Context is used for tracing.
func (c *MemoryClient) CheckCache(ctx context.Context, key string) (bool, []byte) {
	result, success := memoryCache.Load(key)
	if !success || result.(cacheRecord).value == nil || result.(cacheRecord).expiry < time.Now().UnixMilli() {
		memoryCache.Delete(key)
		return false, nil
	}
	return true, result.(cacheRecord).value
}

// AddToCache takes in Context for tracing, key for storing in the cache, a value in the form of a byte array for storing in the cache,
// a TTL for how long from now the record will exist. If the key exists, the value is not nil, and the TTL is not
// expired then the record will be added to the cache.
func (c *MemoryClient) AddToCache(ctx context.Context, key string, value []byte, ttl int64) {
	result, success := memoryCache.Load(key)
	if !success || result.(cacheRecord).value == nil || result.(cacheRecord).expiry < time.Now().UnixMilli() {
		log.Debugf("Adding to Memory Cache: %s", key)
		var record cacheRecord
		record.expiry = ttl + time.Now().UnixMilli()
		record.value = value
		memoryCache.Store(key, record)
	}
}
