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

// AttributesByID returns the character attributes indicated by the id field, the context is
// used for tracing. If the character attributes are cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func AttributesByID(ctx context.Context, id *int) (*model.Attributes, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "CharacterAttributesByID")
	defer span.End()
	var characterAttributes *model.Attributes = new(model.Attributes)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/characters/%s/attributes", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "CharacterAttributesByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return characterAttributes, err
	}

	if err := json.Unmarshal(responseBytes, &characterAttributes); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return characterAttributes, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return characterAttributes, nil
}
