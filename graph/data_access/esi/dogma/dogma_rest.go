package dogma

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	"github.com/cryanbrow/eve-graphql-go/graph/model"
	log "github.com/sirupsen/logrus"
)

func DogmaAttributeByID(id *int) (*model.DogmaAttributeDetail, error) {
	var dogmaAttribute *model.DogmaAttributeDetail = new(model.DogmaAttributeDetail)
	if id == nil {
		return nil, nil
	}
	base_url := fmt.Sprintf("%s/dogma/attributes/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "DogmaAttributeByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
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
		return nil, nil
	}
	base_url := fmt.Sprintf("%s/dogma/effects/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "DogmaEffectByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return dogmaEffect, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaEffect); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaEffect, err
	}

	return dogmaEffect, nil
}