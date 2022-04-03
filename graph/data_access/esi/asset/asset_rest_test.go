package asset

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

type AssetMockMakeCachingRESTCallType func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	AssetMockMakeCachingRESTCall AssetMockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
	return m.AssetMockMakeCachingRESTCall(ctx, baseURL, verb, body, additionalQueryParams, redisQueryKey)
}
