package caching

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type cacheRecord struct {
	value  []byte
	expiry int64
}

func CheckMemoryCache(key string) (bool, []byte) {
	result, success := Cache.Load(key)
	if !success || result.(cacheRecord).value == nil || result.(cacheRecord).expiry < time.Now().UnixMilli() {
		Cache.Delete(key)
		return false, nil
	} else {
		return true, result.(cacheRecord).value
	}
}

func AddToMemoryCache(key string, value []byte, ttl int64) {
	result, success := Cache.Load(key)
	if !success || result.(cacheRecord).value == nil || result.(cacheRecord).expiry < time.Now().UnixMilli() {
		log.Debugf("Adding to Memory Cache: %s", key)
		var record cacheRecord
		record.expiry = ttl + time.Now().UnixMilli()
		record.value = value
		Cache.Store(key, record)
	}
}

var (
	Cache sync.Map
)
