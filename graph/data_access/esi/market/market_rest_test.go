package market

import (
	"bytes"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additionalQueryParams, redisQueryKey)
}
