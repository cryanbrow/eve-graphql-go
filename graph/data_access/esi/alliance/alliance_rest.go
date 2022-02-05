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
	baseUrl := fmt.Sprintf("%s/alliances/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "AllianceByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
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
	MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

func SetupAllianceRest() {
	restHelper = &helpers.RestHelperClient{}
}
