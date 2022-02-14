package universe

import (
	"bytes"
	"context"
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 5

	resp, err := CategoryByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Accessories"
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := CategoryByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallCategoryByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := CategoryByID(context.Background(), &testID)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 5

	_, err := CategoryByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
