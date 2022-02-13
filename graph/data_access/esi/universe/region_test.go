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
*              RegionByID              *
***************************************/

func TestSuccessfulRegionByID(t *testing.T) {
	jsonResponse := `{
		"constellations": [
		  20000096,
		  20000097,
		  20000098,
		  20000099,
		  20000100,
		  20000101,
		  20000102,
		  20000103,
		  20000104,
		  20000105,
		  20000106,
		  20000107
		],
		"description": "Scalding Pass is a vast stellar nursery, the birth place of stars. A harsh and unforgiving region of space home to violent ion storms and rocked by intense solar winds, the Pass is littered with the wrecks of ships that were blindsided by the unpredictable dangers this region holds.\r\n\r\nHowever, there are some who manage to find solace and safety in such a tumultuous place. The Cartel is rumored to have many outlying bases in the region, masked by stellar phenomena, and in recent times many capsuleer groups have braved its burning embrace in pursuit of their dark dreams of power.",
		"name": "Scalding Pass",
		"region_id": 10000008
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 10000008

	resp, err := RegionByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Scalding Pass"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDRegionByID(t *testing.T) {
	var testID *int

	_, err := RegionByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallRegionByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 10000008

	_, err := RegionByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalRegionByID(t *testing.T) {
	jsonResponse := `{{
		"constellations": [
		  20000096,
		  20000097,
		  20000098,
		  20000099,
		  20000100,
		  20000101,
		  20000102,
		  20000103,
		  20000104,
		  20000105,
		  20000106,
		  20000107
		],
		"description": "Scalding Pass is a vast stellar nursery, the birth place of stars. A harsh and unforgiving region of space home to violent ion storms and rocked by intense solar winds, the Pass is littered with the wrecks of ships that were blindsided by the unpredictable dangers this region holds.\r\n\r\nHowever, there are some who manage to find solace and safety in such a tumultuous place. The Cartel is rumored to have many outlying bases in the region, masked by stellar phenomena, and in recent times many capsuleer groups have braved its burning embrace in pursuit of their dark dreams of power.",
		"name": "Scalding Pass",
		"region_id": 10000008
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 10000008

	_, err := RegionByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
