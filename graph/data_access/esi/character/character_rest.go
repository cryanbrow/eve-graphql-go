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
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/characters/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "CharacterByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return character, err
	}

	if err := json.Unmarshal(responseBytes, &character); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return character, err
	}

	return character, nil
}
