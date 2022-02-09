package character

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

func CharacterByID(id *int) (*model.Character, error) {
	var character *model.Character = new(model.Character)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/characters/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "CharacterByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return character, err
	}

	if err := json.Unmarshal(responseBytes, &character); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return character, err
	}

	return character, nil
}

func CorporationHistory(id *int) ([]*model.CorporationHistory, error) {
	var corpHistory []*model.CorporationHistory = make([]*model.CorporationHistory, 0)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/characters/%s/corporationhistory", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "CorporationHistory:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return corpHistory, err
	}

	if err := json.Unmarshal(responseBytes, &corpHistory); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return corpHistory, err
	}

	return corpHistory, nil
}

type RestHelper interface {
	MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

func SetupCharacterRest() {
	restHelper = &helpers.RestHelperClient{}
}
