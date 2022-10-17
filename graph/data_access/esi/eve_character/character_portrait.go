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
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// PortraitByID returns the character portrait indicated by the id field, the context is
// used for tracing. If the character portrait is cached the ESI will not be called until the ttl
// and the cached instance will be returned.
func PortraitByID(ctx context.Context, id *int) (*model.CharacterPortrait, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "CharacterPortraitByID")
	defer span.End()
	var characterPortrait *model.CharacterPortrait = new(model.CharacterPortrait)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/characters/%s/portrait", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "CharacterPortraitByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return characterPortrait, err
	}

	if err := json.Unmarshal(responseBytes, &characterPortrait); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return characterPortrait, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return characterPortrait, nil
}
