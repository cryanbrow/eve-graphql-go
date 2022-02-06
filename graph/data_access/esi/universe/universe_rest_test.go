package universe

import (
	"bytes"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

/***************************************
*             MOCK SECTION             *
***************************************/

type MockAddToRedisCacheType func(key string, value []byte, ttl int64)
type MockCheckRedisCacheType func(key string) (bool, []byte)

type MockRedisClient struct {
	MockAdd   MockAddToRedisCacheType
	MockCheck MockCheckRedisCacheType
}

func (m *MockRedisClient) AddToRedisCache(key string, value []byte, ttl int64) {
	m.MockAdd(key, value, ttl)
}

func (m *MockRedisClient) CheckRedisCache(key string) (bool, []byte) {
	return m.MockCheck(key)
}

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additionalQueryParams, redisQueryKey)
}
