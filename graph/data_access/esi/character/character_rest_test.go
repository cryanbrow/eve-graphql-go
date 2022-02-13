package character

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

type CharacterMockMakeCachingRESTCallType func(ctx context.Context, baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	CharacterMockMakeCachingRESTCall CharacterMockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(ctx context.Context, baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.CharacterMockMakeCachingRESTCall(ctx, baseUrl, verb, body, additionalQueryParams, redisQueryKey)
}
