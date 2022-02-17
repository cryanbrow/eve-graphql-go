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

// StargateDetails takes a context for tracing and an array of IDs to query the Stargates by.
func StargateDetails(ctx context.Context, stargates []*int) ([]*model.Stargate, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "StargateDetails")
	defer span.End()
	stargateDetails := make([]*model.Stargate, 0)
	for _, element := range stargates {
		stargate, err := StargateByID(newCtx, element)
		if err == nil {
			stargateDetails = append(stargateDetails, stargate)
		} else {
			return nil, err
		}
	}
	return stargateDetails, nil
}

// StargateByID takes a context for tracing and an ID to query the Stargate by.
func StargateByID(ctx context.Context, id *int) (*model.Stargate, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "StargateByID")
	defer span.End()
	var stargate *model.Stargate = new(model.Stargate)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/stargates/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "StargateByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return stargate, err
	}

	if err := json.Unmarshal(responseBytes, &stargate); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return stargate, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return stargate, nil
}
