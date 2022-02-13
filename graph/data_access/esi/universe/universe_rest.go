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
	AddToCache(key string, value []byte, ttl int64, ctx context.Context)
	CheckCache(key string, ctx context.Context) (bool, []byte)
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
	MakeCachingRESTCall(ctx context.Context, baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)
}
