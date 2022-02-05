package corporation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"
)

func CorporationByID(id *int) (*model.Corporation, error) {
	var corporation *model.Corporation = new(model.Corporation)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/corporations/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "CorporationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := rest_helper.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return corporation, err
	}

	if err := json.Unmarshal(responseBytes, &corporation); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return corporation, err
	}

	return corporation, nil
}

type RestHelper interface {
	MakeCachingRESTCall(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)
}

var (
	rest_helper RestHelper
)

func SetupCorporationRest() {
	rest_helper = &helpers.RestHelperClient{}
}
