package corporation

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

func TestSuccessful_CorporationByID(t *testing.T) {
	jsonResponse := `{
		"ceo_id": 3018996,
		"creator_id": 1,
		"description": "The Minmatar heart sings for freedom and the Minmatar soul strives for open skies, but the Minmatar heart has been withering away in captivity. It is now up to you, capsuleer. You hold the power to free our people. You are the heroes of your generation. Join us in the struggle for freedom. Death to Amarr; long live the Minmatar Nation.",
		"faction_id": 500002,
		"home_station_id": 60015096,
		"member_count": 22862,
		"name": "Tribal Liberation Force",
		"shares": 0,
		"tax_rate": 0,
		"ticker": "TLIB",
		"url": ""
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	resp, err := CorporationByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Tribal Liberation Force"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_CorporationByID(t *testing.T) {
	jsonResponse := `{
		"ceo_id": 3018996,
		"creator_id": 1,
		"description": "The Minmatar heart sings for freedom and the Minmatar soul strives for open skies, but the Minmatar heart has been withering away in captivity. It is now up to you, capsuleer. You hold the power to free our people. You are the heroes of your generation. Join us in the struggle for freedom. Death to Amarr; long live the Minmatar Nation.",
		"faction_id": 500002,
		"home_station_id": 60015096,
		"member_count": 22862,
		"name": "Tribal Liberation Force",
		"shares": 0,
		"tax_rate": 0,
		"ticker": "TLIB",
		"url": ""
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := CorporationByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_CorporationByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := CorporationByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_CorporationByID(t *testing.T) {
	jsonResponse := `{{
		"ceo_id": 3018996,
		"creator_id": 1,
		"description": "The Minmatar heart sings for freedom and the Minmatar soul strives for open skies, but the Minmatar heart has been withering away in captivity. It is now up to you, capsuleer. You hold the power to free our people. You are the heroes of your generation. Join us in the struggle for freedom. Death to Amarr; long live the Minmatar Nation.",
		"faction_id": 500002,
		"home_station_id": 60015096,
		"member_count": 22862,
		"name": "Tribal Liberation Force",
		"shares": 0,
		"tax_rate": 0,
		"ticker": "TLIB",
		"url": ""
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := CorporationByID(&test_id)
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
