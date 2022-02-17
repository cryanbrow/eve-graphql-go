package corporation

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const tracerName = "github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/corporation"

// ByID takes a context for tracing and an ID to query the corporation by.
func ByID(ctx context.Context, id *int) (*model.Corporation, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "CorporationByID")
	defer span.End()
	var corporation *model.Corporation = new(model.Corporation)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/corporations/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "CorporationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return corporation, err
	}

	if err := json.Unmarshal(responseBytes, &corporation); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return corporation, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return corporation, nil
}

// RestHelper is an interface used in mocking for unit tests.
type RestHelper interface {
	MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

// SetupCorporationRest configures dependencies for the Corporation rest package
func SetupCorporationRest() {
	restHelper = &helpers.RestHelperClient{}
}
