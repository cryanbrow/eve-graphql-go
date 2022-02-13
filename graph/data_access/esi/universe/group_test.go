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

func TestSuccessfulGroupByID(t *testing.T) {
	jsonResponse := `{
		"category_id": 6,
		"group_id": 25,
		"name": "Frigate",
		"published": true,
		"types": [
		  582,
		  583,
		  584,
		  585,
		  586,
		  587,
		  589,
		  590,
		  591,
		  592,
		  593,
		  594,
		  595,
		  597,
		  598,
		  599,
		  600,
		  602,
		  603,
		  605,
		  607,
		  608,
		  609,
		  613,
		  614,
		  616,
		  618,
		  619,
		  1896,
		  1898,
		  1900,
		  1902,
		  2161,
		  3532,
		  3751,
		  3753,
		  3766,
		  3768,
		  11019,
		  11940,
		  11942,
		  17360,
		  17619,
		  17703,
		  17705,
		  17707,
		  17812,
		  17841,
		  17924,
		  17926,
		  17928,
		  17930,
		  17932,
		  29248,
		  32880,
		  32983,
		  32985,
		  32987,
		  32989,
		  33190,
		  33468,
		  33655,
		  33657,
		  33659,
		  33661,
		  33663,
		  33665,
		  33667,
		  33669,
		  33677,
		  33816,
		  34443,
		  37453,
		  37454,
		  37455,
		  37456,
		  47269,
		  54731,
		  58745
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 25

	resp, err := GroupByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Frigate"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDGroupByID(t *testing.T) {
	jsonResponse := `{
		"category_id": 6,
		"group_id": 25,
		"name": "Frigate",
		"published": true,
		"types": [
		  582,
		  583,
		  584,
		  585,
		  586,
		  587,
		  589,
		  590,
		  591,
		  592,
		  593,
		  594,
		  595,
		  597,
		  598,
		  599,
		  600,
		  602,
		  603,
		  605,
		  607,
		  608,
		  609,
		  613,
		  614,
		  616,
		  618,
		  619,
		  1896,
		  1898,
		  1900,
		  1902,
		  2161,
		  3532,
		  3751,
		  3753,
		  3766,
		  3768,
		  11019,
		  11940,
		  11942,
		  17360,
		  17619,
		  17703,
		  17705,
		  17707,
		  17812,
		  17841,
		  17924,
		  17926,
		  17928,
		  17930,
		  17932,
		  29248,
		  32880,
		  32983,
		  32985,
		  32987,
		  32989,
		  33190,
		  33468,
		  33655,
		  33657,
		  33659,
		  33661,
		  33663,
		  33665,
		  33667,
		  33669,
		  33677,
		  33816,
		  34443,
		  37453,
		  37454,
		  37455,
		  37456,
		  47269,
		  54731,
		  58745
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := GroupByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallGroupByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 25

	_, err := GroupByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalGroupByID(t *testing.T) {
	jsonResponse := `{{
		"category_id": 6,
		"group_id": 25,
		"name": "Frigate",
		"published": true,
		"types": [
		  582,
		  583,
		  584,
		  585,
		  586,
		  587,
		  589,
		  590,
		  591,
		  592,
		  593,
		  594,
		  595,
		  597,
		  598,
		  599,
		  600,
		  602,
		  603,
		  605,
		  607,
		  608,
		  609,
		  613,
		  614,
		  616,
		  618,
		  619,
		  1896,
		  1898,
		  1900,
		  1902,
		  2161,
		  3532,
		  3751,
		  3753,
		  3766,
		  3768,
		  11019,
		  11940,
		  11942,
		  17360,
		  17619,
		  17703,
		  17705,
		  17707,
		  17812,
		  17841,
		  17924,
		  17926,
		  17928,
		  17930,
		  17932,
		  29248,
		  32880,
		  32983,
		  32985,
		  32987,
		  32989,
		  33190,
		  33468,
		  33655,
		  33657,
		  33659,
		  33661,
		  33663,
		  33665,
		  33667,
		  33669,
		  33677,
		  33816,
		  34443,
		  37453,
		  37454,
		  37455,
		  37456,
		  47269,
		  54731,
		  58745
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 25

	_, err := GroupByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
