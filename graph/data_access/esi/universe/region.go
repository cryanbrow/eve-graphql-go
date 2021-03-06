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

// RegionByID takes a context for tracing and an ID to query the Region by.
func RegionByID(ctx context.Context, id *int) (*model.Region, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "RegionByID")
	defer span.End()
	var region *model.Region = new(model.Region)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/regions/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "RegionByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return region, err
	}

	if err := json.Unmarshal(responseBytes, &region); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return region, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return region, nil
}
