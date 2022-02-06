package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

const byteArrayFail string = "Failed to return correct byte array."

var jsonResponse string
var testUrl string

func TestSuccessfulMakeCachingRESTCall(t *testing.T) {
	queryParams, shouldReturn := setupSuccessfulRESTCall()
	if shouldReturn {
		return
	}

	var buffer bytes.Buffer
	bytes, _, err := restHelper.MakeCachingRESTCall(testUrl, http.MethodGet, buffer, queryParams, "himom")
	if string(bytes) != jsonResponse {
		t.Error(byteArrayFail)
	}
	if err != nil {
		t.Errorf(ErrorWasNotNil, err)
	}
}

func setupSuccessfulRESTCall() ([]configuration.Key_value, bool) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	setRedisClient()
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	queryParams := make([]configuration.Key_value, 2)
	kv := new(configuration.Key_value)
	kv.Key = "page"
	kv.Value = strconv.Itoa(1)
	queryParams = append(queryParams, *kv)
	return queryParams, false
}

func TestInCacheSuccessfulMakeCachingRESTCall(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	RedisClientVar = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return true, make([]byte, 0)
		},
	}

	queryParams := make([]configuration.Key_value, 2)
	kv := new(configuration.Key_value)
	kv.Key = "page"
	kv.Value = strconv.Itoa(1)
	queryParams = append(queryParams, *kv)

	var buffer bytes.Buffer
	_, _, err := restHelper.MakeCachingRESTCall(testUrl, http.MethodGet, buffer, queryParams, "himom")
	if err != nil {
		t.Errorf(ErrorWasNotNil, err)
	}
}

func TestSuccessfulWithDefaultParamsMakeCachingRESTCall(t *testing.T) {
	queryParams, shouldReturn := setupSuccessfulRESTCall()
	if shouldReturn {
		return
	}

	configuration.AppConfig.Esi.Default.QueryParams = queryParams

	var buffer bytes.Buffer
	bytes4, _, err := restHelper.MakeCachingRESTCall(testUrl, http.MethodGet, buffer, nil, "himom")
	byteString := string(bytes4)
	if byteString != jsonResponse {
		fmt.Printf("expected: %s : actual %s", jsonResponse, byteString)
		t.Error(byteArrayFail)
	}
	if err != nil {
		t.Errorf(ErrorWasNotNil, err)
	}
}

func TestSuccessfulWithQueryParamsMakeCachingRESTCall(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	setRedisClient()
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	var buffer bytes.Buffer
	bytes, _, err := restHelper.MakeCachingRESTCall(testUrl, http.MethodGet, buffer, nil, "himom")
	if string(bytes) != jsonResponse {
		t.Error(byteArrayFail)
	}
	if err != nil {
		t.Errorf(ErrorWasNotNil, err)
	}
}

func TestUnparseableURLMakeCachingRESTCall(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	setRedisClient()
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
		t.Error(NilError)
	}
}

func TestNewRequestFailureMakeCachingRESTCall(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	setRedisClient()
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
		t.Error(NilError)
	}
}

func TestDoFailureMakeCachingRESTCall(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	setRedisClient()
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
		t.Error(NilError)
	}
}

func Test404FailureMakeCachingRESTCall(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	setRedisClient()
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
		t.Error(NilError)
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
	MakeCachingRESTCall(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

func init() {
	restHelper = &RestHelperClient{}
	jsonResponse = `[{
		"full_name": "mock-repo"
	   }]`
	testUrl = "https://www.google.com"
}

func setRedisClient() {
	RedisClientVar = &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, make([]byte, 0)
		},
	}
}
