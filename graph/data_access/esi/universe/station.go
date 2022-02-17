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

// StationsByIDs takes a context for tracing and an array of IDs to query the Stations by.
func StationsByIDs(ctx context.Context, ids []*int) ([]*model.Station, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "StationsByIDs")
	defer span.End()
	stationDetails := make([]*model.Station, 0)
	for _, element := range ids {
		station, err := StationByID(newCtx, element)
		if err == nil {
			stationDetails = append(stationDetails, station)
		} else {
			return nil, err
		}
	}
	return stationDetails, nil
}

// StationByID takes a context for tracing and an ID to query the Station by.
func StationByID(ctx context.Context, id *int) (*model.Station, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "StationByID")
	defer span.End()
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	if *id > 2147483647 {
		return nil, nil
	}
	var station *model.Station = new(model.Station)
	baseURL := fmt.Sprintf("%s/universe/stations/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "StationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return station, err
	}

	if err := json.Unmarshal(responseBytes, &station); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return station, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return station, nil
}
