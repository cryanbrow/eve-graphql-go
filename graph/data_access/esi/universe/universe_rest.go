package universe

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/caching"
	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

const tracer_name = "github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"

type RedisClientInterface interface {
	AddToRedisCache(key string, value []byte, ttl int64)
	CheckRedisCache(key string) (bool, []byte)
}

var (
	RedisClient RedisClientInterface
	restHelper  RestHelper
)

func SetupUniverseRest() {
	RedisClient = &caching.Client{}
	restHelper = &helpers.RestHelperClient{}
}

type RestHelper interface {
	MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error)
}
