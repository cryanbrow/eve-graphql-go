package universe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	model "github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"
)

func GraphicByID(id *int) (*model.Graphic, error) {
	var graphic *model.Graphic = new(model.Graphic)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/universe/graphics/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "GraphicByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return graphic, err
	}

	if err := json.Unmarshal(responseBytes, &graphic); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return graphic, err
	}

	return graphic, nil
}
