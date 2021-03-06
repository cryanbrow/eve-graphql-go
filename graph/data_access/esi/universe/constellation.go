package universe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	model "github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// ConstellationsByIDs takes a context for tracing and an array of IDs to query the Constellations by.
func ConstellationsByIDs(ctx context.Context, ids []*int) ([]*model.Constellation, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "ConstellationsByIDs")
	defer span.End()
	constellationDetails := make([]*model.Constellation, 0)
	for _, element := range ids {
		constellation, err := ConstellationByID(newCtx, element)
		if err == nil {
			constellationDetails = append(constellationDetails, constellation)
		} else {
			return nil, err
		}
	}
	return constellationDetails, nil
}

// ConstellationByID takes a context for tracing and an ID to query the Constellation by.
func ConstellationByID(ctx context.Context, id *int) (*model.Constellation, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "ConstellationByID")
	defer span.End()
	var constellation *model.Constellation = new(model.Constellation)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/constellations/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "ConstellationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return constellation, err
	}

	if err := json.Unmarshal(responseBytes, &constellation); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return constellation, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return constellation, nil
}
