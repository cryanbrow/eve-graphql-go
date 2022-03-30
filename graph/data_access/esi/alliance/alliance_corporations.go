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
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	local_model "github.com/cryanbrow/eve-graphql-go/graph/model"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// CorporationsByName returns the alliance indicated by the name field, the context is
// used for tracing. If the alliance is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func CorporationsByName(ctx context.Context, name *string) (*[]int, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AllianceCorporationsByName")
	defer span.End()
	allianceID, err := universe.IDForName(newCtx, name, local_model.Alliances)
	if err != nil {
		return nil, errors.New("unknown name for alliance")
	}
	return CorporationsByID(newCtx, &allianceID)
}

// CorporationsByID returns the alliance indicated by the id field, the context is
// used for tracing. If the alliance is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func CorporationsByID(ctx context.Context, id *int) (*[]int, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AllianceCorporationsByID")
	defer span.End()
	var corporations = make([]int, 0)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/alliances/%s/corporations", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "AllianceCorporationsByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return &corporations, err
	}

	if err := json.Unmarshal(responseBytes, &corporations); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return &corporations, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return &corporations, nil
}
