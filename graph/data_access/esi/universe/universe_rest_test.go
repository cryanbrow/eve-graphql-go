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

type MockAddToCacheType func(key string, value []byte, ttl int64, ctx context.Context)
type MockCheckCacheType func(key string, ctx context.Context) (bool, []byte)

type MockCachingClient struct {
	MockAdd   MockAddToCacheType
	MockCheck MockCheckCacheType
}

func (m *MockCachingClient) AddToCache(key string, value []byte, ttl int64, ctx context.Context) {
	m.MockAdd(key, value, ttl, ctx)
}

func (m *MockCachingClient) CheckCache(key string, ctx context.Context) (bool, []byte) {
	return m.MockCheck(key, ctx)
}

type MockMakeCachingRESTCallType func(ctx context.Context, baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(ctx context.Context, baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(ctx, baseUrl, verb, body, additionalQueryParams, redisQueryKey)
}
