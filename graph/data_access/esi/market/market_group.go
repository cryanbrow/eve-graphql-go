package market

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

func GroupByID(ctx context.Context, id *int) (*model.MarketGroup, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "MarketGroupByID")
	defer span.End()
	var marketGroup *model.MarketGroup = new(model.MarketGroup)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/markets/groups/%s/", configuration.AppConfig.Esi.Default.URL, strconv.Itoa(*id))
	redisKey := "MarketGroupByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return marketGroup, err
	}

	if err := json.Unmarshal(responseBytes, &marketGroup); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return marketGroup, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return marketGroup, nil
}
