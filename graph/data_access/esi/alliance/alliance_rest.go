package alliance

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	local_model "github.com/cryanbrow/eve-graphql-go/graph/model"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const tracerName = "github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/alliance"

// ByName returns the alliance indicated by the name field, the context is
// used for tracing. If the alliance is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func ByName(ctx context.Context, name *string) (*model.Alliance, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AllianceByName")
	defer span.End()
	allianceID, err := universe.IDForName(newCtx, name, local_model.Alliances)
	if err != nil {
		return nil, errors.New("unknown name for region")
	}
	return ByID(newCtx, &allianceID)
}

// ByID returns the alliance indicated by the id field, the context is
// used for tracing. If the alliance is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func ByID(ctx context.Context, id *int) (*model.Alliance, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AllianceByID")
	defer span.End()
	var alliance *model.Alliance = new(model.Alliance)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/alliances/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "AllianceByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return alliance, err
	}

	if err := json.Unmarshal(responseBytes, &alliance); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return alliance, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return alliance, nil
}

// RestHelper is an interface used in mocking for unit tests.
type RestHelper interface {
	MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

// SetupAllianceRest injects required dependencies into the alliance package.
func SetupAllianceRest() {
	restHelper = &helpers.RestHelperClient{}
}
