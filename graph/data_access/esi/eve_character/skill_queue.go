package eve_character

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

// SkillQueueByName returns the character skill queue indicated by the name field, the context is
// used for tracing. If the character skill queue is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func SkillQueueByName(ctx context.Context, name *string) ([]*model.SkillQueueItem, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "SkillQueueByName")
	defer span.End()
	characterID, err := universe.IDForName(newCtx, name, local_model.Characters)
	if err != nil {
		return nil, errors.New("unknown name for region")
	}
	return SkillQueueByID(newCtx, &characterID)
}

// SkillQueueByID returns the character skill queue indicated by the id field, the context is
// used for tracing. If the character skill queue is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func SkillQueueByID(ctx context.Context, id *int) ([]*model.SkillQueueItem, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "SkillQueueByID")
	defer span.End()
	var skillQueue []*model.SkillQueueItem = make([]*model.SkillQueueItem, 0)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/characters/%s/skillqueue", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "SkillQueueByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return skillQueue, err
	}

	if err := json.Unmarshal(responseBytes, &skillQueue); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return skillQueue, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return skillQueue, nil
}
