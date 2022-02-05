package helpers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

func TestSuccessful_MakeCachingRESTCall(t *testing.T) {
	Redis_client = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, make([]byte, 0)
		},
	}
	jsonResponse := `[{
		"full_name": "mock-repo"
	   }]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	query_params := make([]configuration.Key_value, 2)
	kv := new(configuration.Key_value)
	kv.Key = "page"
	kv.Value = strconv.Itoa(1)
	query_params = append(query_params, *kv)

	url := "https://www.google.com"
	var buffer bytes.Buffer
	bytes, _, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, query_params, "himom")
	if string(bytes) != jsonResponse {
		t.Error("Failed to return correct byte array.")
	}
	if err != nil {
		t.Error("Returned non nil error.")
	}
}

func TestInCacheSuccessful_MakeCachingRESTCall(t *testing.T) {
	jsonResponse := `[{
		"full_name": "mock-repo"
	   }]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	Redis_client = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return true, make([]byte, 0)
		},
	}

	query_params := make([]configuration.Key_value, 2)
	kv := new(configuration.Key_value)
	kv.Key = "page"
	kv.Value = strconv.Itoa(1)
	query_params = append(query_params, *kv)

	url := "https://www.google.com"
	var buffer bytes.Buffer
	_, _, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, query_params, "himom")
	if err != nil {
		t.Error("Returned non nil error.")
	}
}

func TestSuccessfulWithDefaultParams_MakeCachingRESTCall(t *testing.T) {
	Redis_client = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, make([]byte, 0)
		},
	}
	jsonResponse := `[{
		"full_name": "mock-repo"
	   }]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	query_params := make([]configuration.Key_value, 2)
	kv := new(configuration.Key_value)
	kv.Key = "page"
	kv.Value = strconv.Itoa(1)
	query_params = append(query_params, *kv)

	configuration.AppConfig.Esi.Default.Query_params = query_params

	url := "https://www.google.com"
	var buffer bytes.Buffer
	bytes, _, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, nil, "himom")
	if string(bytes) != jsonResponse {
		t.Error("Failed to return correct byte array.")
	}
	if err != nil {
		t.Error("Returned non nil error.")
	}
}

func TestSuccessfulWithQueryParams_MakeCachingRESTCall(t *testing.T) {
	Redis_client = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, make([]byte, 0)
		},
	}
	jsonResponse := `[{
		"full_name": "mock-repo"
	   }]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	url := "https://www.google.com"
	var buffer bytes.Buffer
	bytes, _, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, nil, "himom")
	if string(bytes) != jsonResponse {
		t.Error("Failed to return correct byte array.")
	}
	if err != nil {
		t.Error("Returned non nil error.")
	}
}

func TestUnparseableURL_MakeCachingRESTCall(t *testing.T) {
	Redis_client = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, make([]byte, 0)
		},
	}
	jsonResponse := `[{
		"full_name": "mock-repo"
	   }]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	myslice := make([]byte, 1)
	myslice[0] = 0x7f
	url := string(myslice)
	var buffer bytes.Buffer
	_, _, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, nil, "himom")
	if err == nil {
		t.Error("Returned nil error.")
	}
}

func TestNewRequestFailure_MakeCachingRESTCall(t *testing.T) {
	Redis_client = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, make([]byte, 0)
		},
	}
	jsonResponse := `[{
		"full_name": "mock-repo"
	   }]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	url := ""
	var buffer bytes.Buffer
	_, _, err := restHelper.MakeCachingRESTCall(url, "Ы", buffer, nil, "himom")
	if err == nil {
		t.Error("Returned nil error.")
	}
}

func TestDoFailure_MakeCachingRESTCall(t *testing.T) {
	Redis_client = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, make([]byte, 0)
		},
	}
	jsonResponse := `[{
		"full_name": "mock-repo"
	   }]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, errors.New("general error")
		},
	}

	url := ""
	var buffer bytes.Buffer
	_, _, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, nil, "himom")
	if err == nil {
		t.Error("Returned nil error.")
	}
}

func Test404Failure_MakeCachingRESTCall(t *testing.T) {
	Redis_client = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, make([]byte, 0)
		},
	}
	jsonResponse := `[{
		"full_name": "mock-repo"
	   }]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       r,
			}, nil
		},
	}

	url := ""
	var buffer bytes.Buffer
	_, _, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, nil, "himom")
	if err == nil {
		t.Error("Returned nil error.")
	}
}

type MockDoType func(req *http.Request) (*http.Response, error)
type MockAddToRedisCacheType func(key string, value []byte, ttl int64)
type MockCheckRedisCacheType func(key string) (bool, []byte)

type MockClient struct {
	MockDo MockDoType
}

type MockRedisClient struct {
	MockAdd   MockAddToRedisCacheType
	MockCheck MockCheckRedisCacheType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func (m *MockRedisClient) AddToRedisCache(key string, value []byte, ttl int64) {
	m.MockAdd(key, value, ttl)
}

func (m *MockRedisClient) CheckRedisCache(key string) (bool, []byte) {
	return m.MockCheck(key)
}

type RestHelper interface {
	MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

func init() {
	restHelper = &RestHelperClient{}
}
