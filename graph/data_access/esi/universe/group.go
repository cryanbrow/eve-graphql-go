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

// GroupByID takes a context for tracing and an ID to query the Group by.
func GroupByID(ctx context.Context, id *int) (*model.Group, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "GroupByID")
	defer span.End()
	var group *model.Group = new(model.Group)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/groups/%s/", configuration.AppConfig.Esi.Default.URL, strconv.Itoa(*id))
	redisKey := "GroupByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return group, err
	}

	if err := json.Unmarshal(responseBytes, &group); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return group, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return group, nil
}
