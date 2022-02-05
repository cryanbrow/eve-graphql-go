package character

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessful_CharacterByID(t *testing.T) {
	jsonResponse := `{
		"birthday": "2008-04-01T10:24:00Z",
		"bloodline_id": 3,
		"corporation_id": 1000182,
		"description": "",
		"gender": "male",
		"name": "Gar Orga",
		"race_id": 2,
		"security_status": 0
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	resp, err := CharacterByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Gar Orga"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_CharacterByID(t *testing.T) {
	jsonResponse := `{
		"birthday": "2008-04-01T10:24:00Z",
		"bloodline_id": 3,
		"corporation_id": 1000182,
		"description": "",
		"gender": "male",
		"name": "Gar Orga",
		"race_id": 2,
		"security_status": 0
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := CharacterByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_CharacterByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := CharacterByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_CharacterByID(t *testing.T) {
	jsonResponse := `{{
		"birthday": "2008-04-01T10:24:00Z",
		"bloodline_id": 3,
		"corporation_id": 1000182,
		"description": "",
		"gender": "male",
		"name": "Gar Orga",
		"race_id": 2,
		"security_status": 0
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := CharacterByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additionalQueryParams, redisQueryKey)
}
