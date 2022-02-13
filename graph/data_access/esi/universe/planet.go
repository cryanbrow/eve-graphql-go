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

func PlanetByID(ctx context.Context, id *int) (*model.Planet, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "PlanetByID")
	defer span.End()
	var planet *model.Planet = new(model.Planet)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/planets/%s/", configuration.AppConfig.Esi.Default.URL, strconv.Itoa(*id))
	redisKey := "PlanetByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return planet, err
	}

	if err := json.Unmarshal(responseBytes, &planet); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return planet, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return planet, nil
}
