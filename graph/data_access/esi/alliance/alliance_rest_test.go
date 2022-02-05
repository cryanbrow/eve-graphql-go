package alliance

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessfulAllianceByID(t *testing.T) {
	jsonResponse := `{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	resp, err := AllianceByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Fleet Coordination Coalition"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDAllianceByID(t *testing.T) {
	jsonResponse := `{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := AllianceByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallAllianceByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := AllianceByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalAllianceByID(t *testing.T) {
	jsonResponse := `{{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := AllianceByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additional_query_params, redis_query_key)
}
