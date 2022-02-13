package universe

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/caching"
	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

const tracerName = "github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"

type CacheClientInterface interface {
	AddToCache(ctx context.Context, key string, value []byte, ttl int64)
	CheckCache(ctx context.Context, key string) (bool, []byte)
}

var (
	CachingClient CacheClientInterface
	restHelper    RestHelper
)

func SetupUniverseRest() {
	CachingClient = caching.Cache
	restHelper = &helpers.RestHelperClient{}
}

type RestHelper interface {
	MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error)
}
