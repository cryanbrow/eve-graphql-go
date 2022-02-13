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
*            SystemByID                *
***************************************/

func TestSuccessfulSystemByID(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000308,
		"name": "Ebolfer",
		"planets": [
		  {
			"planet_id": 40133749
		  },
		  {
			"moons": [
			  40133751
			],
			"planet_id": 40133750
		  },
		  {
			"asteroid_belts": [
			  40133753
			],
			"planet_id": 40133752
		  },
		  {
			"moons": [
			  40133755,
			  40133756,
			  40133757,
			  40133758,
			  40133759,
			  40133760,
			  40133761,
			  40133762,
			  40133763,
			  40133764
			],
			"planet_id": 40133754
		  },
		  {
			"asteroid_belts": [
			  40133766,
			  40133767,
			  40133769,
			  40133770,
			  40133775,
			  40133782
			],
			"moons": [
			  40133768,
			  40133771,
			  40133772,
			  40133773,
			  40133774,
			  40133776,
			  40133777,
			  40133778,
			  40133779,
			  40133780,
			  40133781,
			  40133783
			],
			"planet_id": 40133765
		  },
		  {
			"asteroid_belts": [
			  40133785
			],
			"moons": [
			  40133786,
			  40133787,
			  40133788,
			  40133789,
			  40133790
			],
			"planet_id": 40133784
		  }
		],
		"position": {
		  "x": -134671986395386050,
		  "y": 15461957069008490,
		  "z": 25148264854745812
		},
		"security_class": "E",
		"security_status": 0.25291144847869873,
		"star_id": 40133748,
		"stargates": [
		  50012050,
		  50012051,
		  50012052,
		  50012053,
		  50012054
		],
		"stations": [
		  60004537,
		  60015096
		],
		"system_id": 30002094
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 30002094

	resp, err := SystemByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Ebolfer"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDSystemByID(t *testing.T) {
	var testID *int

	_, err := SystemByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallSystemByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 30002094

	_, err := SystemByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalSystemByID(t *testing.T) {
	jsonResponse := `{{
		"constellation_id": 20000308,
		"name": "Ebolfer",
		"planets": [
		  {
			"planet_id": 40133749
		  },
		  {
			"moons": [
			  40133751
			],
			"planet_id": 40133750
		  },
		  {
			"asteroid_belts": [
			  40133753
			],
			"planet_id": 40133752
		  },
		  {
			"moons": [
			  40133755,
			  40133756,
			  40133757,
			  40133758,
			  40133759,
			  40133760,
			  40133761,
			  40133762,
			  40133763,
			  40133764
			],
			"planet_id": 40133754
		  },
		  {
			"asteroid_belts": [
			  40133766,
			  40133767,
			  40133769,
			  40133770,
			  40133775,
			  40133782
			],
			"moons": [
			  40133768,
			  40133771,
			  40133772,
			  40133773,
			  40133774,
			  40133776,
			  40133777,
			  40133778,
			  40133779,
			  40133780,
			  40133781,
			  40133783
			],
			"planet_id": 40133765
		  },
		  {
			"asteroid_belts": [
			  40133785
			],
			"moons": [
			  40133786,
			  40133787,
			  40133788,
			  40133789,
			  40133790
			],
			"planet_id": 40133784
		  }
		],
		"position": {
		  "x": -134671986395386050,
		  "y": 15461957069008490,
		  "z": 25148264854745812
		},
		"security_class": "E",
		"security_status": 0.25291144847869873,
		"star_id": 40133748,
		"stargates": [
		  50012050,
		  50012051,
		  50012052,
		  50012053,
		  50012054
		],
		"stations": [
		  60004537,
		  60015096
		],
		"system_id": 30002094
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 30002094

	_, err := SystemByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

func TestSuccessfulSystemsByIDs(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000308,
		"name": "Ebolfer",
		"planets": [
		  {
			"planet_id": 40133749
		  },
		  {
			"moons": [
			  40133751
			],
			"planet_id": 40133750
		  },
		  {
			"asteroid_belts": [
			  40133753
			],
			"planet_id": 40133752
		  },
		  {
			"moons": [
			  40133755,
			  40133756,
			  40133757,
			  40133758,
			  40133759,
			  40133760,
			  40133761,
			  40133762,
			  40133763,
			  40133764
			],
			"planet_id": 40133754
		  },
		  {
			"asteroid_belts": [
			  40133766,
			  40133767,
			  40133769,
			  40133770,
			  40133775,
			  40133782
			],
			"moons": [
			  40133768,
			  40133771,
			  40133772,
			  40133773,
			  40133774,
			  40133776,
			  40133777,
			  40133778,
			  40133779,
			  40133780,
			  40133781,
			  40133783
			],
			"planet_id": 40133765
		  },
		  {
			"asteroid_belts": [
			  40133785
			],
			"moons": [
			  40133786,
			  40133787,
			  40133788,
			  40133789,
			  40133790
			],
			"planet_id": 40133784
		  }
		],
		"position": {
		  "x": -134671986395386050,
		  "y": 15461957069008490,
		  "z": 25148264854745812
		},
		"security_class": "E",
		"security_status": 0.25291144847869873,
		"star_id": 40133748,
		"stargates": [
		  50012050,
		  50012051,
		  50012052,
		  50012053,
		  50012054
		],
		"stations": [
		  60004537,
		  60015096
		],
		"system_id": 30002094
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID1 int = 30002094
	var testID2 int = 30002094
	var ids []*int = make([]*int, 2)
	ids[0] = &testID1
	ids[1] = &testID2

	resp, err := SystemsByIDs(context.Background(), ids)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Ebolfer"
	if *resp[0].Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailureSystemsByIDs(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID1 int = 30002094
	var testID2 int = 30002094
	var ids []*int = make([]*int, 2)
	ids[0] = &testID1
	ids[1] = &testID2

	_, err := SystemsByIDs(context.Background(), ids)
	if err == nil {
		t.Error(helpers.NilError)
	}
}
