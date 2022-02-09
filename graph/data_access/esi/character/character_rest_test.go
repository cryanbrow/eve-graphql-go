package character

import (
	"bytes"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

type CharacterMockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	CharacterMockMakeCachingRESTCall CharacterMockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.CharacterMockMakeCachingRESTCall(baseUrl, verb, body, additionalQueryParams, redisQueryKey)
}
