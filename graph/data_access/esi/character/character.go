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

// ByID returns the character indicated by the id field, the context is
// used for tracing. If the character is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func ByID(ctx context.Context, id *int) (*model.Character, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "CharacterByID")
	defer span.End()
	var character *model.Character = new(model.Character)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/characters/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "CharacterByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return character, err
	}

	if err := json.Unmarshal(responseBytes, &character); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return character, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return character, nil
}
