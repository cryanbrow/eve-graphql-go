package dogma

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

func DogmaAttributeByID(id *int) (*model.DogmaAttributeDetail, error) {
	var dogmaAttribute *model.DogmaAttributeDetail = new(model.DogmaAttributeDetail)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/dogma/attributes/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "DogmaAttributeByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := rest_helper.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return dogmaAttribute, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaAttribute); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaAttribute, err
	}

	return dogmaAttribute, nil
}

func DogmaEffectByID(id *int) (*model.DogmaEffectDetail, error) {
	var dogmaEffect *model.DogmaEffectDetail = new(model.DogmaEffectDetail)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/dogma/effects/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "DogmaEffectByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := rest_helper.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return dogmaEffect, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaEffect); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaEffect, err
	}

	return dogmaEffect, nil
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
