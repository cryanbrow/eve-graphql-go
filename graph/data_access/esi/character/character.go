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

//CharacterByID returns the character indicated by the id field, the context is
//used for tracing. If the character is cached the ESI will not be called until the ttl
//and the cached instance will be returned.
func CharacterByID(ctx context.Context, id *int) (*model.Character, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "CharacterByID")
	defer span.End()
	var character *model.Character = new(model.Character)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/characters/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "CharacterByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey, newCtx)
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
