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

func ItemTypeByID(ctx context.Context, id *int) (*model.ItemType, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "ItemTypeByID")
	defer span.End()
	var itemType *model.ItemType = new(model.ItemType)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/universe/types/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "ItemTypeByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return itemType, err
	}

	if err := json.Unmarshal(responseBytes, &itemType); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return itemType, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return itemType, nil
}

func ItemTypesByIDs(ctx context.Context, itemTypes []*int) ([]*model.ItemType, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "ItemTypesByIDs")
	defer span.End()
	itemTypeDetails := make([]*model.ItemType, 0)
	for _, element := range itemTypes {
		itemType, err := ItemTypeByID(newCtx, element)
		if err == nil {
			itemTypeDetails = append(itemTypeDetails, itemType)
		} else {
			return nil, err
		}
	}
	return itemTypeDetails, nil
}
