package universe

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

/***************************************
*             MOCK SECTION             *
***************************************/

type MockAddToRedisCacheType func(key string, value []byte, ttl int64, ctx context.Context)
type MockCheckRedisCacheType func(key string, ctx context.Context) (bool, []byte)

type MockRedisClient struct {
	MockAdd   MockAddToRedisCacheType
	MockCheck MockCheckRedisCacheType
}

func (m *MockRedisClient) AddToRedisCache(key string, value []byte, ttl int64, ctx context.Context) {
	m.MockAdd(key, value, ttl, ctx)
}

func (m *MockRedisClient) CheckRedisCache(key string, ctx context.Context) (bool, []byte) {
	return m.MockCheck(key, ctx)
}

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additionalQueryParams, redisQueryKey, ctx)
}
