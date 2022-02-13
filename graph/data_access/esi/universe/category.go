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

func CategoryByID(ctx context.Context, id *int) (*model.Category, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "CategoryByID")
	defer span.End()
	var category *model.Category = new(model.Category)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/universe/categories/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "CategoryByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return category, err
	}

	if err := json.Unmarshal(responseBytes, &category); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return category, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return category, nil
}
