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

func SystemsByIDs(ids []*int, ctx context.Context) ([]*model.System, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "SystemsByIDs")
	defer span.End()
	systemDetails := make([]*model.System, 0)
	for _, element := range ids {
		system, err := SystemByID(element, newCtx)
		if err == nil {
			systemDetails = append(systemDetails, system)
		} else {
			return nil, err
		}
	}
	return systemDetails, nil
}

func SystemByID(id *int, ctx context.Context) (*model.System, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "SystemByID")
	defer span.End()
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	var system *model.System = new(model.System)
	baseUrl := fmt.Sprintf("%s/universe/systems/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "SystemByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey, newCtx)
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
