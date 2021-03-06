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

// MoonDetails takes a context for tracing and an array of IDs to query the Moons by.
func MoonDetails(ctx context.Context, moons []*int) ([]*model.Moon, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "MoonDetails")
	defer span.End()
	moonDetails := make([]*model.Moon, 0)
	for _, element := range moons {
		moon, err := MoonByID(newCtx, element)
		if err == nil {
			moonDetails = append(moonDetails, moon)
		} else {
			return nil, err
		}
	}
	return moonDetails, nil
}

// MoonByID takes a context for tracing and an ID to query the Moon by.
func MoonByID(ctx context.Context, id *int) (*model.Moon, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "MoonByID")
	defer span.End()
	var moon *model.Moon = new(model.Moon)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/moons/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "MoonByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return moon, err
	}

	if err := json.Unmarshal(responseBytes, &moon); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return moon, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return moon, nil
}
