package universe

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

/***************************************
*             GraphicByID              *
***************************************/

func TestSuccessfulGraphicByID(t *testing.T) {
	jsonResponse := `{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 21573

	resp, err := GraphicByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "amarrnavy"
	if *resp.SofFationName != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDGraphicByID(t *testing.T) {
	jsonResponse := `{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := GraphicByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallGraphicByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 21573

	_, err := GraphicByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalGraphicByID(t *testing.T) {
	jsonResponse := `{{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 21573

	_, err := GraphicByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
