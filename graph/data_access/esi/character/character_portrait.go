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

func CharacterPortraitByID(id *int) (*model.CharacterPortrait, error) {
	var characterPortrait *model.CharacterPortrait = new(model.CharacterPortrait)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/characters/%s/portrait", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "CharacterPortraitByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return characterPortrait, err
	}

	if err := json.Unmarshal(responseBytes, &characterPortrait); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return characterPortrait, err
	}

	return characterPortrait, nil
}
