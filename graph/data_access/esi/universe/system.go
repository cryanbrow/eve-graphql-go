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

// SystemsByIDs takes a context for tracing and an array of IDs to query the Systems by.
func SystemsByIDs(ctx context.Context, ids []*int) ([]*model.System, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "SystemsByIDs")
	defer span.End()
	systemDetails := make([]*model.System, 0)
	for _, element := range ids {
		system, err := SystemByID(newCtx, element)
		if err == nil {
			systemDetails = append(systemDetails, system)
		} else {
			return nil, err
		}
	}
	return systemDetails, nil
}

// SystemByID takes a context for tracing and an ID to query the System by.
func SystemByID(ctx context.Context, id *int) (*model.System, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "SystemByID")
	defer span.End()
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	var system *model.System = new(model.System)
	baseURL := fmt.Sprintf("%s/universe/systems/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "SystemByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return system, err
	}

	if err := json.Unmarshal(responseBytes, &system); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return system, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return system, nil
}
