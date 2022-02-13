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

type MockAddToCacheType func(ctx context.Context, key string, value []byte, ttl int64)
type MockCheckCacheType func(ctx context.Context, key string) (bool, []byte)

type MockCachingClient struct {
	MockAdd   MockAddToCacheType
	MockCheck MockCheckCacheType
}

func (m *MockCachingClient) AddToCache(ctx context.Context, key string, value []byte, ttl int64) {
	m.MockAdd(ctx, key, value, ttl)
}

func (m *MockCachingClient) CheckCache(ctx context.Context, key string) (bool, []byte) {
	return m.MockCheck(ctx, key)
}

type MockMakeCachingRESTCallType func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(ctx, baseURL, verb, body, additionalQueryParams, redisQueryKey)
}
