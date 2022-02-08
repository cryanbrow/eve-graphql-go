package market

import (
	"bytes"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

type RestHelper interface {
	MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

func SetupMarketRest() {
	restHelper = &helpers.RestHelperClient{}
}
