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
*            ItemTypeByID              *
***************************************/

func TestSuccessfulItemTypeByID(t *testing.T) {
	jsonResponse := `{
		"capacity": 160,
		"description": "The Kestrel is a heavy missile boat with one of the most sophisticated sensor arrays around. Interestingly enough, it has been used by both the Caldari Navy and several wealthy trade corporations as a cargo-hauling vessel. It is one of few trading vessels with good punching power, making it ideal for solo trade-runs in dangerous areas. The Kestrel was designed so that it could take up to four missile launchers but as a result it can not be equipped with turret weapons nor with mining lasers.",
		"dogma_attributes": [
		  {
			"attribute_id": 3,
			"value": 0
		  },
		  {
			"attribute_id": 4,
			"value": 1113000
		  },
		  {
			"attribute_id": 9,
			"value": 400
		  },
		  {
			"attribute_id": 11,
			"value": 45
		  },
		  {
			"attribute_id": 12,
			"value": 2
		  },
		  {
			"attribute_id": 13,
			"value": 4
		  },
		  {
			"attribute_id": 14,
			"value": 4
		  },
		  {
			"attribute_id": 15,
			"value": 0
		  },
		  {
			"attribute_id": 18,
			"value": 0
		  },
		  {
			"attribute_id": 19,
			"value": 1
		  },
		  {
			"attribute_id": 21,
			"value": 0
		  },
		  {
			"attribute_id": 37,
			"value": 325
		  },
		  {
			"attribute_id": 38,
			"value": 160
		  },
		  {
			"attribute_id": 552,
			"value": 38
		  },
		  {
			"attribute_id": 48,
			"value": 180
		  },
		  {
			"attribute_id": 49,
			"value": 0
		  },
		  {
			"attribute_id": 564,
			"value": 620
		  },
		  {
			"attribute_id": 55,
			"value": 165000
		  },
		  {
			"attribute_id": 2113,
			"value": 1
		  },
		  {
			"attribute_id": 1547,
			"value": 1
		  },
		  {
			"attribute_id": 70,
			"value": 3.27
		  },
		  {
			"attribute_id": 524,
			"value": 0.75
		  },
		  {
			"attribute_id": 588,
			"value": 5
		  },
		  {
			"attribute_id": 525,
			"value": 1
		  },
		  {
			"attribute_id": 600,
			"value": 5
		  },
		  {
			"attribute_id": 101,
			"value": 4
		  },
		  {
			"attribute_id": 102,
			"value": 0
		  },
		  {
			"attribute_id": 1132,
			"value": 400
		  },
		  {
			"attribute_id": 109,
			"value": 0.67
		  },
		  {
			"attribute_id": 110,
			"value": 0.67
		  },
		  {
			"attribute_id": 111,
			"value": 0.67
		  },
		  {
			"attribute_id": 1137,
			"value": 3
		  },
		  {
			"attribute_id": 1555,
			"value": 25
		  },
		  {
			"attribute_id": 633,
			"value": 0
		  },
		  {
			"attribute_id": 124,
			"value": 16777215
		  },
		  {
			"attribute_id": 129,
			"value": 6
		  },
		  {
			"attribute_id": 1154,
			"value": 3
		  },
		  {
			"attribute_id": 136,
			"value": 1
		  },
		  {
			"attribute_id": 661,
			"value": 3000
		  },
		  {
			"attribute_id": 662,
			"value": 0.05
		  },
		  {
			"attribute_id": 153,
			"value": 0.00000224
		  },
		  {
			"attribute_id": 1178,
			"value": 100
		  },
		  {
			"attribute_id": 1179,
			"value": 0.01
		  },
		  {
			"attribute_id": 161,
			"value": 19700
		  },
		  {
			"attribute_id": 162,
			"value": 45.98
		  },
		  {
			"attribute_id": 113,
			"value": 0.67
		  },
		  {
			"attribute_id": 1196,
			"value": 0.01
		  },
		  {
			"attribute_id": 1198,
			"value": 0.01
		  },
		  {
			"attribute_id": 1199,
			"value": 100
		  },
		  {
			"attribute_id": 1200,
			"value": 100
		  },
		  {
			"attribute_id": 182,
			"value": 3330
		  },
		  {
			"attribute_id": 192,
			"value": 5
		  },
		  {
			"attribute_id": 1224,
			"value": 1
		  },
		  {
			"attribute_id": 208,
			"value": 0
		  },
		  {
			"attribute_id": 209,
			"value": 0
		  },
		  {
			"attribute_id": 210,
			"value": 0
		  },
		  {
			"attribute_id": 211,
			"value": 11
		  },
		  {
			"attribute_id": 217,
			"value": 395
		  },
		  {
			"attribute_id": 1768,
			"value": 11324
		  },
		  {
			"attribute_id": 1259,
			"value": 0.63
		  },
		  {
			"attribute_id": 1261,
			"value": 0.63
		  },
		  {
			"attribute_id": 1262,
			"value": 0.25
		  },
		  {
			"attribute_id": 246,
			"value": 395
		  },
		  {
			"attribute_id": 1271,
			"value": 0
		  },
		  {
			"attribute_id": 1281,
			"value": 1
		  },
		  {
			"attribute_id": 263,
			"value": 500
		  },
		  {
			"attribute_id": 264,
			"value": 0
		  },
		  {
			"attribute_id": 265,
			"value": 350
		  },
		  {
			"attribute_id": 267,
			"value": 0.5
		  },
		  {
			"attribute_id": 268,
			"value": 0.9
		  },
		  {
			"attribute_id": 269,
			"value": 0.75
		  },
		  {
			"attribute_id": 270,
			"value": 0.55
		  },
		  {
			"attribute_id": 271,
			"value": 1
		  },
		  {
			"attribute_id": 272,
			"value": 0.5
		  },
		  {
			"attribute_id": 273,
			"value": 0.6
		  },
		  {
			"attribute_id": 274,
			"value": 0.8
		  },
		  {
			"attribute_id": 277,
			"value": 1
		  },
		  {
			"attribute_id": 283,
			"value": 0
		  },
		  {
			"attribute_id": 422,
			"value": 1
		  },
		  {
			"attribute_id": 76,
			"value": 50000
		  },
		  {
			"attribute_id": 463,
			"value": 10
		  },
		  {
			"attribute_id": 79,
			"value": 4000
		  },
		  {
			"attribute_id": 479,
			"value": 625000
		  },
		  {
			"attribute_id": 482,
			"value": 330
		  },
		  {
			"attribute_id": 484,
			"value": 0.75
		  },
		  {
			"attribute_id": 2045,
			"value": 1
		  },
		  {
			"attribute_id": 511,
			"value": 0
		  }
		],
		"dogma_effects": [
		  {
			"effect_id": 5080,
			"is_default": false
		  },
		  {
			"effect_id": 5234,
			"is_default": false
		  },
		  {
			"effect_id": 5237,
			"is_default": false
		  },
		  {
			"effect_id": 5240,
			"is_default": false
		  },
		  {
			"effect_id": 5243,
			"is_default": false
		  }
		],
		"graphic_id": 313,
		"group_id": 25,
		"market_group_id": 61,
		"mass": 1113000,
		"name": "Kestrel",
		"packaged_volume": 2500,
		"portion_size": 1,
		"published": true,
		"radius": 45.98,
		"type_id": 602,
		"volume": 19700
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 602

	resp, err := ItemTypeByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Kestrel"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDItemTypeByID(t *testing.T) {
	var testID *int

	_, err := ItemTypeByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallItemTypeByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 602

	_, err := ItemTypeByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalItemTypeByID(t *testing.T) {
	jsonResponse := `{{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 602

	_, err := ItemTypeByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

func TestSuccessfulItemTypesByIDs(t *testing.T) {
	jsonResponse := `{
		"capacity": 160,
		"description": "The Kestrel is a heavy missile boat with one of the most sophisticated sensor arrays around. Interestingly enough, it has been used by both the Caldari Navy and several wealthy trade corporations as a cargo-hauling vessel. It is one of few trading vessels with good punching power, making it ideal for solo trade-runs in dangerous areas. The Kestrel was designed so that it could take up to four missile launchers but as a result it can not be equipped with turret weapons nor with mining lasers.",
		"dogma_attributes": [
		  {
			"attribute_id": 3,
			"value": 0
		  },
		  {
			"attribute_id": 4,
			"value": 1113000
		  },
		  {
			"attribute_id": 9,
			"value": 400
		  },
		  {
			"attribute_id": 11,
			"value": 45
		  },
		  {
			"attribute_id": 12,
			"value": 2
		  },
		  {
			"attribute_id": 13,
			"value": 4
		  },
		  {
			"attribute_id": 14,
			"value": 4
		  },
		  {
			"attribute_id": 15,
			"value": 0
		  },
		  {
			"attribute_id": 18,
			"value": 0
		  },
		  {
			"attribute_id": 19,
			"value": 1
		  },
		  {
			"attribute_id": 21,
			"value": 0
		  },
		  {
			"attribute_id": 37,
			"value": 325
		  },
		  {
			"attribute_id": 38,
			"value": 160
		  },
		  {
			"attribute_id": 552,
			"value": 38
		  },
		  {
			"attribute_id": 48,
			"value": 180
		  },
		  {
			"attribute_id": 49,
			"value": 0
		  },
		  {
			"attribute_id": 564,
			"value": 620
		  },
		  {
			"attribute_id": 55,
			"value": 165000
		  },
		  {
			"attribute_id": 2113,
			"value": 1
		  },
		  {
			"attribute_id": 1547,
			"value": 1
		  },
		  {
			"attribute_id": 70,
			"value": 3.27
		  },
		  {
			"attribute_id": 524,
			"value": 0.75
		  },
		  {
			"attribute_id": 588,
			"value": 5
		  },
		  {
			"attribute_id": 525,
			"value": 1
		  },
		  {
			"attribute_id": 600,
			"value": 5
		  },
		  {
			"attribute_id": 101,
			"value": 4
		  },
		  {
			"attribute_id": 102,
			"value": 0
		  },
		  {
			"attribute_id": 1132,
			"value": 400
		  },
		  {
			"attribute_id": 109,
			"value": 0.67
		  },
		  {
			"attribute_id": 110,
			"value": 0.67
		  },
		  {
			"attribute_id": 111,
			"value": 0.67
		  },
		  {
			"attribute_id": 1137,
			"value": 3
		  },
		  {
			"attribute_id": 1555,
			"value": 25
		  },
		  {
			"attribute_id": 633,
			"value": 0
		  },
		  {
			"attribute_id": 124,
			"value": 16777215
		  },
		  {
			"attribute_id": 129,
			"value": 6
		  },
		  {
			"attribute_id": 1154,
			"value": 3
		  },
		  {
			"attribute_id": 136,
			"value": 1
		  },
		  {
			"attribute_id": 661,
			"value": 3000
		  },
		  {
			"attribute_id": 662,
			"value": 0.05
		  },
		  {
			"attribute_id": 153,
			"value": 0.00000224
		  },
		  {
			"attribute_id": 1178,
			"value": 100
		  },
		  {
			"attribute_id": 1179,
			"value": 0.01
		  },
		  {
			"attribute_id": 161,
			"value": 19700
		  },
		  {
			"attribute_id": 162,
			"value": 45.98
		  },
		  {
			"attribute_id": 113,
			"value": 0.67
		  },
		  {
			"attribute_id": 1196,
			"value": 0.01
		  },
		  {
			"attribute_id": 1198,
			"value": 0.01
		  },
		  {
			"attribute_id": 1199,
			"value": 100
		  },
		  {
			"attribute_id": 1200,
			"value": 100
		  },
		  {
			"attribute_id": 182,
			"value": 3330
		  },
		  {
			"attribute_id": 192,
			"value": 5
		  },
		  {
			"attribute_id": 1224,
			"value": 1
		  },
		  {
			"attribute_id": 208,
			"value": 0
		  },
		  {
			"attribute_id": 209,
			"value": 0
		  },
		  {
			"attribute_id": 210,
			"value": 0
		  },
		  {
			"attribute_id": 211,
			"value": 11
		  },
		  {
			"attribute_id": 217,
			"value": 395
		  },
		  {
			"attribute_id": 1768,
			"value": 11324
		  },
		  {
			"attribute_id": 1259,
			"value": 0.63
		  },
		  {
			"attribute_id": 1261,
			"value": 0.63
		  },
		  {
			"attribute_id": 1262,
			"value": 0.25
		  },
		  {
			"attribute_id": 246,
			"value": 395
		  },
		  {
			"attribute_id": 1271,
			"value": 0
		  },
		  {
			"attribute_id": 1281,
			"value": 1
		  },
		  {
			"attribute_id": 263,
			"value": 500
		  },
		  {
			"attribute_id": 264,
			"value": 0
		  },
		  {
			"attribute_id": 265,
			"value": 350
		  },
		  {
			"attribute_id": 267,
			"value": 0.5
		  },
		  {
			"attribute_id": 268,
			"value": 0.9
		  },
		  {
			"attribute_id": 269,
			"value": 0.75
		  },
		  {
			"attribute_id": 270,
			"value": 0.55
		  },
		  {
			"attribute_id": 271,
			"value": 1
		  },
		  {
			"attribute_id": 272,
			"value": 0.5
		  },
		  {
			"attribute_id": 273,
			"value": 0.6
		  },
		  {
			"attribute_id": 274,
			"value": 0.8
		  },
		  {
			"attribute_id": 277,
			"value": 1
		  },
		  {
			"attribute_id": 283,
			"value": 0
		  },
		  {
			"attribute_id": 422,
			"value": 1
		  },
		  {
			"attribute_id": 76,
			"value": 50000
		  },
		  {
			"attribute_id": 463,
			"value": 10
		  },
		  {
			"attribute_id": 79,
			"value": 4000
		  },
		  {
			"attribute_id": 479,
			"value": 625000
		  },
		  {
			"attribute_id": 482,
			"value": 330
		  },
		  {
			"attribute_id": 484,
			"value": 0.75
		  },
		  {
			"attribute_id": 2045,
			"value": 1
		  },
		  {
			"attribute_id": 511,
			"value": 0
		  }
		],
		"dogma_effects": [
		  {
			"effect_id": 5080,
			"is_default": false
		  },
		  {
			"effect_id": 5234,
			"is_default": false
		  },
		  {
			"effect_id": 5237,
			"is_default": false
		  },
		  {
			"effect_id": 5240,
			"is_default": false
		  },
		  {
			"effect_id": 5243,
			"is_default": false
		  }
		],
		"graphic_id": 313,
		"group_id": 25,
		"market_group_id": 61,
		"mass": 1113000,
		"name": "Kestrel",
		"packaged_volume": 2500,
		"portion_size": 1,
		"published": true,
		"radius": 45.98,
		"type_id": 602,
		"volume": 19700
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID1 int = 602
	var testID2 int = 602
	var ids []*int = make([]*int, 2)
	ids[0] = &testID1
	ids[1] = &testID2

	resp, err := ItemTypesByIDs(context.Background(), ids)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Kestrel"
	if *resp[0].Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailureItemTypesByIDs(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID1 int = 602
	var testID2 int = 602
	var ids []*int = make([]*int, 2)
	ids[0] = &testID1
	ids[1] = &testID2

	_, err := ItemTypesByIDs(context.Background(), ids)
	if err == nil {
		t.Error(helpers.NilError)
	}
}
