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

// IconByName returns the alliance icon indicated by the name field, the context is
// used for tracing. If the alliance is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func IconByName(ctx context.Context, name *string) (*model.Icon, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AllianceIconByName")
	defer span.End()
	allianceID, err := universe.IDForName(newCtx, name, local_model.Alliances)
	if err != nil {
		return nil, errors.New("unknown name for alliance")
	}
	return IconByID(newCtx, &allianceID)
}

// CorporationsByID returns the alliance icon indicated by the id field, the context is
// used for tracing. If the alliance is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func IconByID(ctx context.Context, id *int) (*model.Icon, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AllianceIconByID")
	defer span.End()
	var icon = model.Icon{}
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/alliances/%s/icons", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "AllianceIconByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return &icon, err
	}

	if err := json.Unmarshal(responseBytes, &icon); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return &icon, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return &icon, nil
}
