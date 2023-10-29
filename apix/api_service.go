package apix

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/sivaosorg/govm/coltx"
	"github.com/sivaosorg/govm/restify"
)

type ApiService interface {
	Do(client *restify.Client, endpoint Endpoint) (*restify.Response, error)
	DoAsync(client *restify.Client, endpoints ...Endpoint)
	DoAsyncWith(client *restify.Client, endpoints map[string]Endpoint)
}

type apiServiceImpl struct {
	conf ApiRequest
}

func NewApiService(conf ApiRequest) ApiService {
	s := &apiServiceImpl{
		conf: conf,
	}
	return s
}

func (s *apiServiceImpl) Do(client *restify.Client, endpoint Endpoint) (*restify.Response, error) {
	return s.execute(client, endpoint)
}

func (s *apiServiceImpl) DoAsync(client *restify.Client, endpoints ...Endpoint) {
	if len(endpoints) == 0 {
		return
	}
	var wg sync.WaitGroup
	for _, v := range endpoints {
		wg.Add(1)
		go func(endpoint Endpoint) {
			defer wg.Done()
			response, err := s.Do(client, endpoint)
			if err != nil {
				// send notification
			} else {
				if response.IsSuccess() {

				}
				if response.IsError() {

				}
			}
		}(v)
	}
	wg.Wait()
}

func (s *apiServiceImpl) DoAsyncWith(client *restify.Client, endpoints map[string]Endpoint) {
	if len(endpoints) == 0 {
		return
	}
	var e []Endpoint
	for _, v := range endpoints {
		if !v.IsEnabled {
			continue
		}
		e = append(e, v)
	}
	s.DoAsync(client, e...)
}

func (s *apiServiceImpl) execute(client *restify.Client, endpoint Endpoint) (*restify.Response, error) {
	if !endpoint.IsEnabled {
		return nil, fmt.Errorf("Endpoint %v unavailable", endpoint.Method)
	}
	fullURL, e := s.conf.CombineUrl(endpoint)
	if e != nil {
		return nil, e
	}
	if client == nil {
		client = restify.New()
	}
	var response *restify.Response
	var err error

	host := s.conf.CombineHostURL(endpoint)
	headers := s.conf.CombineHeaders(endpoint)
	retry := s.conf.CombineRetry(endpoint)
	auth := s.conf.CombineAuthentication(endpoint)

	client.SetHostURL(host)
	client.SetHeaders(headers)
	client.SetDebug(endpoint.DebugMode)

	if endpoint.AvailableTimeout() {
		client.SetTimeout(endpoint.Timeout)
	}
	if endpoint.AvailableQueryParams() {
		client.SetQueryParams(endpoint.QueryParams)
	}
	if endpoint.AvailablePathParams() {
		client.SetPathParams(endpoint.PathParams)
	}
	if auth.IsEnabled {
		if strings.EqualFold("basic", auth.Type) {
			client.SetBasicAuth(auth.Username, auth.Password)
		}
		if strings.EqualFold("token", auth.Type) {
			client.SetHeader("Authorization", auth.Token)
		}
	}
	if retry.IsEnabled {
		retryFunc := func(_response *restify.Response, err error) bool {
			if _response == nil {
				return false
			}
			if retry.AvailableRetryOnStatus() {
				confirm := coltx.Contains(retry.RetryOnStatus, _response.StatusCode()) || err != nil
				if confirm {
					// send notification
					// use goroutine
				}
				return confirm
			} else {
				confirm := (_response.StatusCode() >= http.StatusBadRequest && _response.StatusCode() <= http.StatusNetworkAuthenticationRequired) || err != nil
				if confirm {
					// send notification
					// use goroutine
				}
				return confirm
			}
		}
		retryAfterFunc := func(c *restify.Client, r *restify.Response) (time.Duration, error) {
			return time.Duration(retry.InitialInterval * time.Duration(retry.BackoffFactor)), nil
		}
		client.
			SetRetryCount(retry.MaxAttempts).
			SetRetryWaitTime(retry.InitialInterval).
			SetRetryMaxWaitTime(retry.MaxInterval).
			AddRetryCondition(retryFunc).
			SetRetryAfter(retryAfterFunc)
	}

	request := client.R()
	request.SetResult(&response).SetError(&err)

	if endpoint.AvailableBody() {
		request.SetBody(endpoint.Body)
	}
	response, err = request.Execute(endpoint.Method, fullURL)
	return response, err
}
