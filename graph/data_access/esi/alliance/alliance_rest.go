package alliance

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

func AllianceByID(id *int) (*model.Alliance, error) {
	var alliance *model.Alliance = new(model.Alliance)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/alliances/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "AllianceByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := rest_helper.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return alliance, err
	}

	if err := json.Unmarshal(responseBytes, &alliance); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return alliance, err
	}

	return alliance, nil
}

type RestHelper interface {
	MakeCachingRESTCall(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)
}

var (
	rest_helper RestHelper
)

func SetupDogmaRest() {
	rest_helper = &helpers.RestHelperClient{}
}
