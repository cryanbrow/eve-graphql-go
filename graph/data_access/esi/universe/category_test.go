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
*             CategoryByID             *
***************************************/

func TestSuccessfulCategoryByID(t *testing.T) {
	jsonResponse := `{
		"category_id": 5,
		"groups": [
		  23,
		  24,
		  872,
		  876,
		  943,
		  1301,
		  1311,
		  1739,
		  1875
		],
		"name": "Accessories",
		"published": true
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 5

	resp, err := CategoryByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Accessories"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDCategoryByID(t *testing.T) {
	jsonResponse := `{
		"category_id": 5,
		"groups": [
		  23,
		  24,
		  872,
		  876,
		  943,
		  1301,
		  1311,
		  1739,
		  1875
		],
		"name": "Accessories",
		"published": true
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := CategoryByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallCategoryByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := CategoryByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalCategoryByID(t *testing.T) {
	jsonResponse := `{{
		"category_id": 5,
		"groups": [
		  23,
		  24,
		  872,
		  876,
		  943,
		  1301,
		  1311,
		  1739,
		  1875
		],
		"name": "Accessories",
		"published": true
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 5

	_, err := CategoryByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
