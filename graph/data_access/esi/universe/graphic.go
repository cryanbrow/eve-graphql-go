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

func GraphicByID(ctx context.Context, id *int) (*model.Graphic, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "GraphicByID")
	defer span.End()
	var graphic *model.Graphic = new(model.Graphic)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/graphics/%s/", configuration.AppConfig.Esi.Default.URL, strconv.Itoa(*id))
	redisKey := "GraphicByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return graphic, err
	}

	if err := json.Unmarshal(responseBytes, &graphic); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return graphic, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return graphic, nil
}
