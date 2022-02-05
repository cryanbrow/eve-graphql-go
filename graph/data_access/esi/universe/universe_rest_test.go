package universe

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

/***************************************
*             AncestryByID             *
***************************************/

func TestSuccessfulInCache_AncestryByID(t *testing.T) {
	jsonResponse := `{
		"bloodline_id": 7,
		"description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
		"icon_id": 1653,
		"id": 13,
		"name": "Activists",
		"short_description": "Making the universe a better place, one fight at a time."
	  }`
	b := []byte(jsonResponse)

	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	Redis_client = mock_redis_client

	var test_id int = 1
	resp, err := AncestryByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Activists"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestSuccessfulNotInCache_AncestryByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
	{
	  "bloodline_id": 7,
	  "description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
	  "icon_id": 1653,
	  "id": 13,
	  "name": "Activists",
	  "short_description": "Making the universe a better place, one fight at a time."
	}
  ]`
	b := []byte(ancestriesJsonResponse)
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	Redis_client = mock_redis_client
	rest_helper = mock_rest_helper

	var test_id int = 13
	resp, err := AncestryByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Activists"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilID_AncestryByID(t *testing.T) {
	var test_id *int = nil
	_, err := AncestryByID(test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailUnmarshalInCache_AncestryByID(t *testing.T) {
	jsonResponse := `{{
		"bloodline_id": 7,
		"description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
		"icon_id": 1653,
		"id": 13,
		"name": "Activists",
		"short_description": "Making the universe a better place, one fight at a time."
	  }`
	b := []byte(jsonResponse)

	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	Redis_client = mock_redis_client

	var test_id int = 13
	_, err := AncestryByID(&test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailUnmarshalNotInCache_AncestryByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
	{{
	  "bloodline_id": 7,
	  "description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
	  "icon_id": 1653,
	  "id": 13,
	  "name": "Activists",
	  "short_description": "Making the universe a better place, one fight at a time."
	}
  ]`
	b := []byte(ancestriesJsonResponse)
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	Redis_client = mock_redis_client

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 13
	_, err := AncestryByID(&test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailRestNotInCache_AncestryByID(t *testing.T) {
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}

	Redis_client = mock_redis_client
	rest_helper = mock_rest_helper

	var test_id int = 13
	_, err := AncestryByID(&test_id)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*          AsteroidBeltDetails         *
***************************************/

func TestSuccessful_AsteroidBeltDetails(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = new(int)
	*test_id = 1
	var ids []*int = make([]*int, 1)
	ids[0] = test_id

	resp, err := AsteroidBeltDetails(ids)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Inaro IX - Asteroid Belt 1"
	if *resp[0].Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilID_AsteroidBeltDetails(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = new(int)
	*test_id = 1
	var ids []*int = make([]*int, 2)
	ids[0] = test_id

	_, err := AsteroidBeltDetails(ids)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*           AsteroidBeltByID           *
***************************************/

func TestSuccessful_AsteroidBeltByID(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	resp, err := AsteroidBeltByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Inaro IX - Asteroid Belt 1"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_AsteroidBeltByID(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := AsteroidBeltByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_AsteroidBeltByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := AsteroidBeltByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_AsteroidBeltByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := AsteroidBeltByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	}

}

/***************************************
*          BloodlineBeltByID           *
***************************************/

func TestSuccessfulInCache_BloodlineByID(t *testing.T) {
	jsonResponse := `{
		"bloodline_id": 5,
		"charisma": 3,
		"corporation_id": 1000066,
		"description": "True Amarrians are proud and supercilious, with a great sense of tradition and ancestry. They are considered arrogant and tyrannical by most others. The Empire's defeat at the hands of the mysterious Jovians, and the Minmatar uprising that followed, left an indelible mark on Amarrian culture. This double failure, a turning point in their history, has shaped an entire generation of policy and philosophy among the imperial elite.",
		"intelligence": 7,
		"memory": 6,
		"name": "Amarr",
		"perception": 4,
		"race_id": 4,
		"ship_type_id": 596,
		"willpower": 10
	  }`
	b := []byte(jsonResponse)

	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	Redis_client = mock_redis_client

	var test_id int = 1
	resp, err := BloodlineByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Amarr"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestSuccessfulNotInCache_BloodlineByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
		{
			"bloodline_id": 5,
			"charisma": 3,
			"corporation_id": 1000066,
			"description": "True Amarrians are proud and supercilious, with a great sense of tradition and ancestry. They are considered arrogant and tyrannical by most others. The Empire's defeat at the hands of the mysterious Jovians, and the Minmatar uprising that followed, left an indelible mark on Amarrian culture. This double failure, a turning point in their history, has shaped an entire generation of policy and philosophy among the imperial elite.",
			"intelligence": 7,
			"memory": 6,
			"name": "Amarr",
			"perception": 4,
			"race_id": 4,
			"ship_type_id": 596,
			"willpower": 10
		  }
  ]`
	b := []byte(ancestriesJsonResponse)
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	Redis_client = mock_redis_client
	rest_helper = mock_rest_helper

	var test_id int = 5
	resp, err := BloodlineByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Amarr"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilID_BloodlineByID(t *testing.T) {
	var test_id *int = nil
	_, err := AncestryByID(test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailUnmarshalInCache_BloodlineByID(t *testing.T) {
	jsonResponse := `{{
		"bloodline_id": 5,
		"charisma": 3,
		"corporation_id": 1000066,
		"description": "True Amarrians are proud and supercilious, with a great sense of tradition and ancestry. They are considered arrogant and tyrannical by most others. The Empire's defeat at the hands of the mysterious Jovians, and the Minmatar uprising that followed, left an indelible mark on Amarrian culture. This double failure, a turning point in their history, has shaped an entire generation of policy and philosophy among the imperial elite.",
		"intelligence": 7,
		"memory": 6,
		"name": "Amarr",
		"perception": 4,
		"race_id": 4,
		"ship_type_id": 596,
		"willpower": 10
	  }`
	b := []byte(jsonResponse)

	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	Redis_client = mock_redis_client

	var test_id int = 5
	_, err := BloodlineByID(&test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailUnmarshalNotInCache_BloodlineByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
	{{
		"bloodline_id": 5,
		"charisma": 3,
		"corporation_id": 1000066,
		"description": "True Amarrians are proud and supercilious, with a great sense of tradition and ancestry. They are considered arrogant and tyrannical by most others. The Empire's defeat at the hands of the mysterious Jovians, and the Minmatar uprising that followed, left an indelible mark on Amarrian culture. This double failure, a turning point in their history, has shaped an entire generation of policy and philosophy among the imperial elite.",
		"intelligence": 7,
		"memory": 6,
		"name": "Amarr",
		"perception": 4,
		"race_id": 4,
		"ship_type_id": 596,
		"willpower": 10
	  }
  ]`
	b := []byte(ancestriesJsonResponse)
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	Redis_client = mock_redis_client

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 5
	_, err := BloodlineByID(&test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailRestNotInCache_BloodlineByID(t *testing.T) {
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}

	Redis_client = mock_redis_client
	rest_helper = mock_rest_helper

	var test_id int = 5
	_, err := BloodlineByID(&test_id)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*             CategoryByID             *
***************************************/

func TestSuccessful_CategoryByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 5

	resp, err := CategoryByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Accessories"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_CategoryByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := CategoryByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_CategoryByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := CategoryByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_CategoryByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 5

	_, err := CategoryByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	}

}

/***************************************
*          ConstellationsByIDs         *
***************************************/

func TestSuccessful_ConstellationsByIDs(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = new(int)
	*test_id = 20000019
	var ids []*int = make([]*int, 1)
	ids[0] = test_id

	resp, err := ConstellationsByIDs(ids)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Ihilakken"
	if *resp[0].Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilID_ConstellationsByIDs(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = new(int)
	*test_id = 20000019
	var ids []*int = make([]*int, 2)
	ids[0] = test_id

	_, err := ConstellationsByIDs(ids)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*          ConstellationByID           *
***************************************/

func TestSuccessful_ConstellationByID(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 20000019

	resp, err := ConstellationByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Ihilakken"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_ConstellationByID(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := ConstellationByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_ConstellationByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 20000019

	_, err := ConstellationByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_ConstellationByID(t *testing.T) {
	jsonResponse := `{{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 20000019

	_, err := ConstellationByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	}

}

/***************************************
*             FactionByID              *
***************************************/

func TestSuccessfulInCache_FactionByID(t *testing.T) {
	jsonResponse := `{
		"corporation_id": 1000084,
		"description": "The largest of the five main empires, the Amarr Empire is a sprawling patch-work of feudal-like provinces held together by the might of the emperor. Religion has always played a big part in Amarrian politics and the Amarrians believe they are the rightful masters of the world, souring their relations with their neighbours. Another source of ill-feelings on part of the other empires is the fact that the Amarrians embrace slavery.",
		"faction_id": 500003,
		"is_unique": true,
		"militia_corporation_id": 1000179,
		"name": "Amarr Empire",
		"size_factor": 5,
		"solar_system_id": 30002187,
		"station_count": 1031,
		"station_system_count": 508
	  }`
	b := []byte(jsonResponse)

	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	Redis_client = mock_redis_client

	var test_id int = 500003
	resp, err := FactionByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Amarr Empire"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestSuccessfulNotInCache_FactionByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
		{
			"corporation_id": 1000084,
			"description": "The largest of the five main empires, the Amarr Empire is a sprawling patch-work of feudal-like provinces held together by the might of the emperor. Religion has always played a big part in Amarrian politics and the Amarrians believe they are the rightful masters of the world, souring their relations with their neighbours. Another source of ill-feelings on part of the other empires is the fact that the Amarrians embrace slavery.",
			"faction_id": 500003,
			"is_unique": true,
			"militia_corporation_id": 1000179,
			"name": "Amarr Empire",
			"size_factor": 5,
			"solar_system_id": 30002187,
			"station_count": 1031,
			"station_system_count": 508
		  }
  ]`
	b := []byte(ancestriesJsonResponse)
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	Redis_client = mock_redis_client
	rest_helper = mock_rest_helper

	var test_id int = 500003
	resp, err := FactionByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Amarr Empire"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilID_FactionByID(t *testing.T) {
	var test_id *int = nil
	_, err := FactionByID(test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailUnmarshalInCache_FactionByID(t *testing.T) {
	jsonResponse := `{{
		"corporation_id": 1000084,
		"description": "The largest of the five main empires, the Amarr Empire is a sprawling patch-work of feudal-like provinces held together by the might of the emperor. Religion has always played a big part in Amarrian politics and the Amarrians believe they are the rightful masters of the world, souring their relations with their neighbours. Another source of ill-feelings on part of the other empires is the fact that the Amarrians embrace slavery.",
		"faction_id": 500003,
		"is_unique": true,
		"militia_corporation_id": 1000179,
		"name": "Amarr Empire",
		"size_factor": 5,
		"solar_system_id": 30002187,
		"station_count": 1031,
		"station_system_count": 508
	  }`
	b := []byte(jsonResponse)

	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	Redis_client = mock_redis_client

	var test_id int = 500003
	_, err := FactionByID(&test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailUnmarshalNotInCache_FactionByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
	{{
		"corporation_id": 1000084,
		"description": "The largest of the five main empires, the Amarr Empire is a sprawling patch-work of feudal-like provinces held together by the might of the emperor. Religion has always played a big part in Amarrian politics and the Amarrians believe they are the rightful masters of the world, souring their relations with their neighbours. Another source of ill-feelings on part of the other empires is the fact that the Amarrians embrace slavery.",
		"faction_id": 500003,
		"is_unique": true,
		"militia_corporation_id": 1000179,
		"name": "Amarr Empire",
		"size_factor": 5,
		"solar_system_id": 30002187,
		"station_count": 1031,
		"station_system_count": 508
	  }
  ]`
	b := []byte(ancestriesJsonResponse)
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	Redis_client = mock_redis_client

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 500003
	_, err := FactionByID(&test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestFailRestNotInCache_FactionByID(t *testing.T) {
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}

	Redis_client = mock_redis_client
	rest_helper = mock_rest_helper

	var test_id int = 500003
	_, err := FactionByID(&test_id)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*             GraphicByID              *
***************************************/

func TestSuccessful_GraphicByID(t *testing.T) {
	jsonResponse := `{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 21573

	resp, err := GraphicByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "amarrnavy"
	if *resp.SofFationName != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_GraphicByID(t *testing.T) {
	jsonResponse := `{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := GraphicByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_GraphicByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 21573

	_, err := GraphicByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_GraphicByID(t *testing.T) {
	jsonResponse := `{{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 21573

	_, err := GraphicByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	}

}

/***************************************
*              GroupByID               *
***************************************/

func TestSuccessful_GroupByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 25

	resp, err := GroupByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Frigate"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_GroupByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := GroupByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_GroupByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 25

	_, err := GroupByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_GroupByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 25

	_, err := GroupByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	}

}

/***************************************
*             MOCK SECTION             *
***************************************/

type MockAddToRedisCacheType func(key string, value []byte, ttl int64)
type MockCheckRedisCacheType func(key string) (bool, []byte)

type MockRedisClient struct {
	MockAdd   MockAddToRedisCacheType
	MockCheck MockCheckRedisCacheType
}

func (m *MockRedisClient) AddToRedisCache(key string, value []byte, ttl int64) {
	m.MockAdd(key, value, ttl)
}

func (m *MockRedisClient) CheckRedisCache(key string) (bool, []byte) {
	return m.MockCheck(key)
}

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additional_query_params, redis_query_key)
}
