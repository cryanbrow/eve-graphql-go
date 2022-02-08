package market

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

func OrderHistory(regionID *int, typeID *int) ([]*model.OrderHistory, error) {
	if regionID == nil || typeID == nil {
		return nil, errors.New(helpers.NilId)
	}
	var orderHistory []*model.OrderHistory = make([]*model.OrderHistory, 0)
	baseUrl := fmt.Sprintf("%s/markets/%s/history", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*regionID))
	redisKey := "OrderHistoryByID:" + strconv.Itoa(*regionID) + ":" + strconv.Itoa(*typeID)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return orderHistory, err
	}

	if err := json.Unmarshal(responseBytes, &orderHistory); err != nil {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return orderHistory, err
	}

	return orderHistory, nil
}
