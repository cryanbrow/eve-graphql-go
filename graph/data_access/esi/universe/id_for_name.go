package universe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	local_model "github.com/cryanbrow/eve-graphql-go/graph/model"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func IDForName(ctx context.Context, name *string, nameType string) (int, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "IDForName")
	defer span.End()
	var ids *local_model.Names = new(local_model.Names)
	baseURL := fmt.Sprintf("%s/universe/ids/", configuration.AppConfig.Esi.Default.URL)
	if name == nil {
		return 0, errors.New("nil name")
	}
	redisKey := "IDForName:" + *name
	singleItemArray := []string{*name}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(singleItemArray)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodPost, buf, nil, redisKey)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(responseBytes, &ids); err != nil {
		log.WithFields(log.Fields{"name": *name}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return 0, err
	}

	span.SetAttributes(attribute.String("request.name", *name), attribute.String("request.nameType", nameType))

	switch nameType {
	case local_model.AGENTS:
		return *ids.Agents[0].ID, nil
	case local_model.ALLIANCES:
		return *ids.Alliances[0].ID, nil
	case local_model.CHARACTERS:
		return *ids.Characters[0].ID, nil
	case local_model.CONSTELLATIONS:
		return *ids.Constellations[0].ID, nil
	case local_model.CORPORATIONS:
		return *ids.Corporations[0].ID, nil
	case local_model.FACTIONS:
		return *ids.Factions[0].ID, nil
	case local_model.INVENTORY_TYPES:
		return *ids.InventoryTypes[0].ID, nil
	case local_model.REGIONS:
		return *ids.Regions[0].ID, nil
	case local_model.SYSTEMS:
		return *ids.Systems[0].ID, nil
	default:
		return 0, errors.New("all fields nil")
	}

}
