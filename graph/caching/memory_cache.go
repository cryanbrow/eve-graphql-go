package caching

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type MemoryClient struct {
}

type cacheRecord struct {
	value  []byte
	expiry int64
}

var memoryCache sync.Map

func (c *MemoryClient) CheckCache(ctx context.Context, key string) (bool, []byte) {
	result, success := memoryCache.Load(key)
	if !success || result.(cacheRecord).value == nil || result.(cacheRecord).expiry < time.Now().UnixMilli() {
		memoryCache.Delete(key)
		return false, nil
	}
	return true, result.(cacheRecord).value
}

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
