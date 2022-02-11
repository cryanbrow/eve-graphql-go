package market

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

const tracer_name = "github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/market"

type RestHelper interface {
	MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

func SetupMarketRest() {
	restHelper = &helpers.RestHelperClient{}
}
