package apix

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/sivaosorg/govm/coltx"
	"github.com/sivaosorg/govm/curlx"
	"github.com/sivaosorg/govm/utils"
)

func NewAuthentication() *Authentication {
	a := &Authentication{}
	return a
}

func (a *Authentication) SetEnabled(value bool) *Authentication {
	a.IsEnabled = value
	return a
}

func (a *Authentication) SetType(value string) *Authentication {
	a.Type = value
	return a
}

func (a *Authentication) SetToken(value string) *Authentication {
	a.Token = value
	return a
}

func (a *Authentication) SetUsername(value string) *Authentication {
	a.Username = value
	return a
}

func (a *Authentication) SetPassword(value string) *Authentication {
	a.Password = value
	return a
}

func (a *Authentication) Json() string {
	return utils.ToJson(a)
}

func GetAuthenticationSample() *Authentication {
	a := NewAuthentication().
		SetEnabled(true).
		SetPassword("pwd").
		SetUsername("user").
		SetType("basic").
		SetToken("<token-here>")
	return a
}

func NewRetry() *Retry {
	r := &Retry{}
	return r
}

func (r *Retry) SetEnabled(value bool) *Retry {
	r.IsEnabled = value
	return r
}

func (r *Retry) SetMaxAttempts(value int) *Retry {
	if value <= 0 {
		log.Panicf("Invalid max_attempts: %v", value)
	}
	r.MaxAttempts = value
	return r
}

func (r *Retry) SetInitialInterval(value time.Duration) *Retry {
	r.InitialInterval = value
	return r
}

func (r *Retry) SetMaxInterval(value time.Duration) *Retry {
	r.MaxInterval = value
	return r
}

func (r *Retry) SetBackoffFactor(value int) *Retry {
	if value < 0 {
		log.Panicf("Invalid backoff_factor: %v", value)
	}
	r.BackoffFactor = value
	return r
}

func (r *Retry) SetRetryOnStatus(values []int) *Retry {
	r.RetryOnStatus = values
	return r
}

func (r *Retry) AppendRetryOnStatus(values ...int) *Retry {
	r.RetryOnStatus = append(r.RetryOnStatus, values...)
	return r
}

func (r *Retry) Json() string {
	return utils.ToJson(r)
}

func RetryValidator(r *Retry) {
	r.SetMaxAttempts(r.MaxAttempts).
		SetBackoffFactor(r.BackoffFactor)
}

func GetRetrySample() *Retry {
	r := NewRetry().
		SetEnabled(true).
		SetBackoffFactor(2).
		SetMaxAttempts(2).
		SetInitialInterval(2*time.Second).
		SetMaxInterval(10*time.Second).
		AppendRetryOnStatus(http.StatusInternalServerError, http.StatusGatewayTimeout)
	return r
}

func NewEndpoint() *Endpoint {
	e := &Endpoint{}
	e.SetRepeat(1)
	return e
}

func (e *Endpoint) SetEnabled(value bool) *Endpoint {
	e.IsEnabled = value
	return e
}

func (e *Endpoint) SetDebugMode(value bool) *Endpoint {
	e.DebugMode = value
	return e
}

func (e *Endpoint) SetBaseURL(value string) *Endpoint {
	u, err := url.Parse(value)
	if err != nil {
		log.Panicf("Invalid base_url: %v", err.Error())
	}
	e.BaseURL = u.String()
	return e
}

func (e *Endpoint) SetTimeout(value time.Duration) *Endpoint {
	e.Timeout = value
	return e
}

func (e *Endpoint) SetPath(value string) *Endpoint {
	e.Path = value
	return e
}

func (e *Endpoint) SetMethod(value string) *Endpoint {
	v, ok := curlx.MethodWithRequestBody[curlx.Method(value)]
	if !ok || !v {
		log.Panicf("Invalid method: %v", value)
	}
	e.Method = value
	return e
}

func (e *Endpoint) SetDescription(value string) *Endpoint {
	e.Description = value
	return e
}

func (e *Endpoint) SetQueryParams(value map[string]string) *Endpoint {
	e.QueryParams = value
	return e
}

func (e *Endpoint) AppendQueryParam(key string, value string) *Endpoint {
	if len(e.QueryParams) == 0 {
		e.SetQueryParams(make(map[string]string))
	}
	e.QueryParams[key] = value
	return e
}

func (e *Endpoint) SetPathParams(value map[string]string) *Endpoint {
	e.PathParams = value
	return e
}

func (e *Endpoint) AppendPathParam(key, value string) *Endpoint {
	if len(e.PathParams) == 0 {
		e.SetPathParams(make(map[string]string))
	}
	e.PathParams[key] = value
	return e
}

func (e *Endpoint) SetHeaders(value map[string]string) *Endpoint {
	e.Headers = value
	return e
}

func (e *Endpoint) AppendHeader(key, value string) *Endpoint {
	if len(e.Headers) == 0 {
		e.SetHeaders(make(map[string]string))
	}
	e.Headers[key] = value
	return e
}

func (e *Endpoint) SetBody(value map[string]interface{}) *Endpoint {
	e.Body = value
	return e
}

func (e *Endpoint) AppendBody(key, value string) *Endpoint {
	if len(e.Body) == 0 {
		e.SetBody(make(map[string]interface{}))
	}
	e.Body[key] = value
	return e
}

func (e *Endpoint) AppendBodyWith(key string, value interface{}) *Endpoint {
	if len(e.Body) == 0 {
		e.SetBody(make(map[string]interface{}))
	}
	e.Body[key] = value
	return e
}

func (e *Endpoint) SetRetry(value Retry) *Endpoint {
	e.Retry = value
	return e
}

func (e *Endpoint) SetAuthentication(value Authentication) *Endpoint {
	e.Authentication = value
	return e
}

func (e *Endpoint) SetRepeat(value int) *Endpoint {
	if value <= 0 {
		log.Panicf("Invalid repeat: %v", value)
	}
	e.Repeat = value
	return e
}

func (e *Endpoint) Json() string {
	return utils.ToJson(e)
}

func EndpointValidator(e *Endpoint) {
	e.SetBaseURL(e.BaseURL).
		SetRepeat(e.Repeat).
		SetMethod(e.Method)
}

func GetEndpointSample() *Endpoint {
	e := NewEndpoint().
		SetEnabled(true).
		SetDebugMode(true).
		SetBaseURL("http://127.0.0.1:8080").
		SetTimeout(10*time.Second).
		SetPath("/api/v1/users").
		SetMethod("POST").
		SetDescription("Create new user").
		AppendHeader("Content-Type", "application/json").
		SetRepeat(2).
		AppendBody("username", "tester").
		AppendBody("email", "tester@gmail.com").
		SetAuthentication(*GetAuthenticationSample()).
		SetRetry(*GetRetrySample())
	return e
}

func NewApiRequest() *ApiRequest {
	a := &ApiRequest{}
	return a
}

func (a *ApiRequest) SetBaseURL(value string) *ApiRequest {
	u, err := url.Parse(value)
	if err != nil {
		log.Panicf("Invalid base_url: %v", err.Error())
	}
	a.BaseURL = u.String()
	return a
}

func (a *ApiRequest) SetAuthentication(value Authentication) *ApiRequest {
	a.Authentication = value
	return a
}

func (a *ApiRequest) SetHeaders(value map[string]string) *ApiRequest {
	a.Headers = value
	return a
}

func (a *ApiRequest) AppendHeader(key, value string) *ApiRequest {
	if len(a.Headers) == 0 {
		a.SetHeaders(make(map[string]string))
	}
	a.Headers[key] = value
	return a
}

func (a *ApiRequest) SetRetry(value Retry) *ApiRequest {
	a.Retry = value
	return a
}

func (a *ApiRequest) SetEndpoints(value map[string]Endpoint) *ApiRequest {
	a.Endpoints = value
	return a
}

func (a *ApiRequest) AppendEndpoint(key string, endpoint Endpoint) *ApiRequest {
	if len(a.Endpoints) == 0 {
		a.SetEndpoints(make(map[string]Endpoint))
	}
	a.Endpoints[key] = endpoint
	return a
}

func (a *ApiRequest) Json() string {
	return utils.ToJson(a)
}

func GetApiRequestSample() *ApiRequest {
	a := NewApiRequest().
		SetBaseURL("http:127.0.0.1:8080").
		SetAuthentication(*GetAuthenticationSample()).
		AppendHeader("Content-Type", "application/json").
		AppendEndpoint("a_endpoint", *GetEndpointSample()).
		AppendEndpoint("b_endpoint", *GetEndpointSample()).
		SetRetry(*GetRetrySample())
	return a
}

func (a *ApiRequest) AvailableEndpoint() bool {
	return len(a.Endpoints) > 0
}

func (a *ApiRequest) AvailableHeader() bool {
	return len(a.Headers) > 0
}

func (a *ApiRequest) GetEndpoint(key string) (Endpoint, error) {
	v, ok := a.Endpoints[key]
	if !ok {
		return Endpoint{}, fmt.Errorf("Endpoint not found")
	}
	return v, nil
}

func (e *Endpoint) Url() (string, error) {
	return e.UrlWith(e.BaseURL)
}

func (e *Endpoint) AvailableBody() bool {
	return len(e.Body) > 0
}

func (e *Endpoint) AvailableQueryParams() bool {
	return len(e.QueryParams) > 0
}

func (e *Endpoint) AvailablePathParams() bool {
	return len(e.PathParams) > 0
}

func (e *Endpoint) AvailableHeaders() bool {
	return len(e.Headers) > 0
}

func (e *Endpoint) UrlWith(baseURL string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	u.Path = e.Path
	query := url.Values{}
	for key, value := range e.QueryParams {
		query.Add(key, value)
	}
	u.RawQuery = query.Encode()
	return u.String(), nil
}

func (e *Endpoint) AvailableTimeout() bool {
	return isDuration(e.Timeout)
}

func (r *Retry) AvailableRetryOnStatus() bool {
	return len(r.RetryOnStatus) > 0
}

func (a *ApiRequest) FilterActivateEndpoints() *ApiRequest {
	endpoints := make(map[string]Endpoint)
	for k, v := range a.Endpoints {
		if !v.IsEnabled {
			continue
		}
		endpoints[k] = v
	}
	a.SetEndpoints(endpoints)
	return a
}

func (a *ApiRequest) CombineHeaders(e Endpoint) map[string]string {
	if a.AvailableHeader() {
		return coltx.MergeMapsString(a.Headers, e.Headers)
	}
	return e.Headers
}

func (a *ApiRequest) CombineAuthentication(e Endpoint) Authentication {
	if a.Authentication.IsEnabled {
		return a.Authentication
	}
	return e.Authentication
}

func (a *ApiRequest) CombineRetry(e Endpoint) Retry {
	if a.Retry.IsEnabled {
		return a.Retry
	}
	return e.Retry
}

func (a *ApiRequest) CombineHostURL(e Endpoint) string {
	if !utils.IsEmpty(a.BaseURL) {
		return a.BaseURL
	}
	return e.BaseURL
}

func (a *ApiRequest) CombineUrl(e Endpoint) (string, error) {
	u, err := e.Url()
	if utils.IsEmpty(a.BaseURL) {
		return u, err
	}
	if err != nil {
		return u, err
	}
	return e.UrlWith(a.BaseURL)
}

func isDuration(t time.Duration) bool {
	return t != 0 && t > 0
}
