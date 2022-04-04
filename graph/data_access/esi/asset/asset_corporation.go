package asset

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	local_model "github.com/cryanbrow/eve-graphql-go/graph/model"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// AssetsByCorporationName returns the corporation assets indicated by the name field, the context is
// used for tracing. If the corporation assets are cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func AssetsByCorporationName(ctx context.Context, name *string) ([]*model.Asset, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AssetsByCorporationName")
	defer span.End()
	corporationID, err := universe.IDForName(newCtx, name, local_model.Corporations)
	if err != nil {
		return nil, errors.New("unknown name for corporation")
	}
	return AssetsByCharacterID(newCtx, &corporationID)
}

// AssetsByCorporationID returns the corporation assets indicated by the id field, the context is
// used for tracing. If the corporation assets are cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func AssetsByCorporationID(ctx context.Context, id *int) ([]*model.Asset, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AssetsByCorporationID")
	defer span.End()
	var assets = make([]*model.Asset, 0)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/corporations/%s/assets", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "AssetsByCorporationID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return assets, err
	}

	if err := json.Unmarshal(responseBytes, &assets); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return assets, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return assets, nil
}
