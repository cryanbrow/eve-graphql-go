package market

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

const tracerName = "github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/market"

// RestHelper is an interface used in mocking for unit tests.
type RestHelper interface {
	MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

func SetupMarketRest() {
	restHelper = &helpers.RestHelperClient{}
}
