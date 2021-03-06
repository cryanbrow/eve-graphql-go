package dogma

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

const tracerName = "github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/dogma"

// AttributeByID takes a context for tracing and an ID to query the Dogma Attribute by.
func AttributeByID(ctx context.Context, id *int) (*model.DogmaAttributeDetail, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "DogmaAttributeByID")
	defer span.End()
	var dogmaAttribute *model.DogmaAttributeDetail = new(model.DogmaAttributeDetail)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/dogma/attributes/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "DogmaAttributeByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return dogmaAttribute, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaAttribute); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaAttribute, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return dogmaAttribute, nil
}

// EffectByID takes a context for tracing and an ID to query the Dogma Effect by.
func EffectByID(ctx context.Context, id *int) (*model.DogmaEffectDetail, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "DogmaEffectByID")
	defer span.End()
	var dogmaEffect *model.DogmaEffectDetail = new(model.DogmaEffectDetail)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/dogma/effects/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "DogmaEffectByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return dogmaEffect, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaEffect); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaEffect, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return dogmaEffect, nil
}

// RestHelper is an interface used in mocking for unit tests.
type RestHelper interface {
	MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

// SetupDogmaRest configures dependencies for the Dogma rest package
func SetupDogmaRest() {
	restHelper = &helpers.RestHelperClient{}
}
