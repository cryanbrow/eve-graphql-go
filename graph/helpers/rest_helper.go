package helpers

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cryanbrow/eve-graphql-go/graph/caching"
	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type RestHelperClient struct {
}

const tracerName = "github.com/cryanbrow/eve-graphql-go/graph/helpers"

func (r *RestHelperClient) MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "MakeCachingRESTCall")
	span.SetAttributes(attribute.String("baseURL", baseURL), attribute.String("verb", verb), attribute.String("redisKey", redisQueryKey))
	defer span.End()
	inCache, result := CachingClientVar.CheckCache(redisQueryKey, newCtx)
	if !inCache {
		crest_url, err := url.Parse(baseURL)
		if err != nil {
			log.WithFields(log.Fields{"baseURL": baseURL, "verb": verb}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, nil, err
		}
		queryParameters := crest_url.Query()
		for _, kv := range configuration.AppConfig.Esi.Default.QueryParams {
			queryParameters.Add(kv.Key, kv.Value)
		}
		for _, kv := range additionalQueryParams {
			queryParameters.Add(kv.Key, kv.Value)
		}

		crest_url.RawQuery = queryParameters.Encode()
		url := crest_url.String()

		log.WithFields(log.Fields{"url": url}).Info("Making REST Call")
		request, err := http.NewRequest(verb, url, &body)
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not build request. : %v", err)
			return make([]byte, 0), nil, err
		}
		response, err := Client.Do(request)
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not make request. : %v", err)
			return make([]byte, 0), nil, err
		}

		h := response.Header
		responseBytes, err := ioutil.ReadAll(response.Body)
		if response.StatusCode != 200 {
			log.WithFields(log.Fields{"url": url, "status_code": response.StatusCode}).Errorf("Received bad status code. : %v", err)
			return make([]byte, 0), nil, errors.New(response.Status)
		}
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not read response for body. : %v", err)
			return make([]byte, 0), nil, err
		}
		CachingClientVar.AddToCache(redisQueryKey, responseBytes, EsiTtlToMillis(h.Get("expires"), newCtx), newCtx)
		return responseBytes, h, nil
	}
	return result, nil, nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type CachingClient interface {
	AddToCache(key string, value []byte, ttl int64, ctx context.Context)
	CheckCache(key string, ctx context.Context) (bool, []byte)
}

var (
	Client           HTTPClient
	CachingClientVar CachingClient
)

func SetupRestHelper() {
	Client = &http.Client{}
	CachingClientVar = caching.Cache
}
