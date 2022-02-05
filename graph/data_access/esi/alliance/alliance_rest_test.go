package alliance

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

func TestSuccessful_AllianceByID(t *testing.T) {
	jsonResponse := `{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	resp, err := AllianceByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Fleet Coordination Coalition"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_AllianceByID(t *testing.T) {
	jsonResponse := `{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := AllianceByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_AllianceByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := AllianceByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_AllianceByID(t *testing.T) {
	jsonResponse := `{{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := AllianceByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	}

}

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additional_query_params, redis_query_key)
}
