package data_access

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type cacheRecord struct {
	value  []byte
	expiry int64
}

func CheckCache(key string) (bool, []byte) {
	result, success := Cache.Load(key)
	if !success || result.(cacheRecord).value == nil || result.(cacheRecord).expiry < time.Now().UnixMilli() {
		Cache.Delete(key)
		return false, nil
	} else {
		return true, result.(cacheRecord).value
	}
}

func AddToCache(key string, value []byte, expiry int64) {
	result, success := Cache.Load(key)
	if !success || result.(cacheRecord).value == nil || result.(cacheRecord).expiry < time.Now().UnixMilli() {
		log.Debugln("Adding to Cache", key)
		var record cacheRecord
		record.expiry = expiry
		record.value = value
		Cache.Store(key, record)
	}
}

var (
	Cache sync.Map
)
