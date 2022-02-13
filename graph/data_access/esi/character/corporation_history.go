package character

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

//CorporationHistory returns the corporation history of the character indicated by the id field, the context is
//used for tracing. If the corporation history is cached the ESI will not be called until the ttl
//and the cached instance will be returned.
func CorporationHistory(ctx context.Context, id *int) ([]*model.CorporationHistory, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "CorporationHistory")
	defer span.End()
	var corpHistory []*model.CorporationHistory = make([]*model.CorporationHistory, 0)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/characters/%s/corporationhistory", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "CorporationHistory:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return corpHistory, err
	}

	if err := json.Unmarshal(responseBytes, &corpHistory); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return corpHistory, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return corpHistory, nil
}
