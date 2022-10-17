package eve_character

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessfulSkillQueueByID(t *testing.T) {
	jsonResponse := `[
		{
		  "finish_date": "2022-10-02T15:29:58Z",
		  "finished_level": 5,
		  "level_end_sp": 512000,
		  "level_start_sp": 90510,
		  "queue_position": 0,
		  "skill_id": 3310,
		  "start_date": "2022-10-01T17:17:50Z",
		  "training_start_sp": 462045
		},
		{
		  "finish_date": "2022-10-10T10:48:27Z",
		  "finished_level": 5,
		  "level_end_sp": 512000,
		  "level_start_sp": 90510,
		  "queue_position": 1,
		  "skill_id": 3311,
		  "start_date": "2022-10-02T15:29:58Z",
		  "training_start_sp": 90557
		}
	]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	resp, err := SkillQueueByID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	if len(resp) != 2 {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDSkillQueueByID(t *testing.T) {
	jsonResponse := `[
		{
		  "finish_date": "2022-10-02T15:29:58Z",
		  "finished_level": 5,
		  "level_end_sp": 512000,
		  "level_start_sp": 90510,
		  "queue_position": 0,
		  "skill_id": 3310,
		  "start_date": "2022-10-01T17:17:50Z",
		  "training_start_sp": 462045
		},
		{
		  "finish_date": "2022-10-10T10:48:27Z",
		  "finished_level": 5,
		  "level_end_sp": 512000,
		  "level_start_sp": 90510,
		  "queue_position": 1,
		  "skill_id": 3311,
		  "start_date": "2022-10-02T15:29:58Z",
		  "training_start_sp": 90557
		}
	]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := SkillQueueByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallSkillQueueByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := SkillQueueByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalSkillQueueByID(t *testing.T) {
	jsonResponse := `{[
		{
		  "finish_date": "2022-10-02T15:29:58Z",
		  "finished_level": 5,
		  "level_end_sp": 512000,
		  "level_start_sp": 90510,
		  "queue_position": 0,
		  "skill_id": 3310,
		  "start_date": "2022-10-01T17:17:50Z",
		  "training_start_sp": 462045
		},
		{
		  "finish_date": "2022-10-10T10:48:27Z",
		  "finished_level": 5,
		  "level_end_sp": 512000,
		  "level_start_sp": 90510,
		  "queue_position": 1,
		  "skill_id": 3311,
		  "start_date": "2022-10-02T15:29:58Z",
		  "training_start_sp": 90557
		}
	]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := SkillQueueByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
