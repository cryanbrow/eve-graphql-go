package character

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

const tracerName = "github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/corporation"

type RestHelper interface {
	MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

//Injects required dependencies into the character package.
func SetupCharacterRest() {
	restHelper = &helpers.RestHelperClient{}
}
