package apix

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/sivaosorg/govm/bot/telegram"
	"github.com/sivaosorg/govm/coltx"
	"github.com/sivaosorg/govm/curlx"
	"github.com/sivaosorg/govm/utils"
)

func NewAuthentication() *AuthenticationConfig {
	a := &AuthenticationConfig{}
	return a
}

func (a *AuthenticationConfig) SetEnabled(value bool) *AuthenticationConfig {
	a.IsEnabled = value
	return a
}

func (a *AuthenticationConfig) SetType(value string) *AuthenticationConfig {
	a.Type = value
	return a
}

func (a *AuthenticationConfig) SetToken(value string) *AuthenticationConfig {
	a.Token = value
	return a
}

func (a *AuthenticationConfig) SetUsername(value string) *AuthenticationConfig {
	a.Username = value
	return a
}

func (a *AuthenticationConfig) SetPassword(value string) *AuthenticationConfig {
	a.Password = value
	return a
}

func (a *AuthenticationConfig) Json() string {
	return utils.ToJson(a)
}

func GetAuthenticationSample() *AuthenticationConfig {
	a := NewAuthentication().
		SetEnabled(false).
		SetPassword("pwd").
		SetUsername("user").
		SetType("basic").
		SetToken("<token-here>")
	return a
}

func NewRetry() *RetryConfig {
	r := &RetryConfig{}
	return r
}

func (r *RetryConfig) SetEnabled(value bool) *RetryConfig {
	r.IsEnabled = value
	return r
}

func (r *RetryConfig) SetMaxAttempts(value int) *RetryConfig {
	if value <= 0 {
		log.Panicf("Invalid max_attempts: %v", value)
	}
	r.MaxAttempts = value
	return r
}

func (r *RetryConfig) SetInitialInterval(value time.Duration) *RetryConfig {
	r.InitialInterval = value
	return r
}

func (r *RetryConfig) SetMaxInterval(value time.Duration) *RetryConfig {
	r.MaxInterval = value
	return r
}

func (r *RetryConfig) SetBackoffFactor(value int) *RetryConfig {
	if value < 0 {
		log.Panicf("Invalid backoff_factor: %v", value)
	}
	r.BackoffFactor = value
	return r
}

func (r *RetryConfig) SetRetryOnStatus(values []int) *RetryConfig {
	r.RetryOnStatus = values
	return r
}

func (r *RetryConfig) AppendRetryOnStatus(values ...int) *RetryConfig {
	r.RetryOnStatus = append(r.RetryOnStatus, values...)
	return r
}

func (r *RetryConfig) Json() string {
	return utils.ToJson(r)
}

func RetryValidator(r *RetryConfig) {
	r.SetMaxAttempts(r.MaxAttempts).
		SetBackoffFactor(r.BackoffFactor)
}

func GetRetrySample() *RetryConfig {
	r := NewRetry().
		SetEnabled(true).
		SetBackoffFactor(2).
		SetMaxAttempts(2).
		SetInitialInterval(2*time.Second).
		SetMaxInterval(10*time.Second).
		AppendRetryOnStatus(http.StatusInternalServerError, http.StatusGatewayTimeout)
	return r
}

func NewEndpoint() *EndpointConfig {
	e := &EndpointConfig{}
	return e
}

func (e *EndpointConfig) SetEnabled(value bool) *EndpointConfig {
	e.IsEnabled = value
	return e
}

func (e *EndpointConfig) SetDebugMode(value bool) *EndpointConfig {
	e.DebugMode = value
	return e
}

func (e *EndpointConfig) SetBaseURL(value string) *EndpointConfig {
	u, err := url.Parse(value)
	if err != nil {
		log.Panicf("Invalid base_url: %v", err.Error())
	}
	e.BaseURL = u.String()
	return e
}

func (e *EndpointConfig) SetTimeout(value time.Duration) *EndpointConfig {
	e.Timeout = value
	return e
}

func (e *EndpointConfig) SetPath(value string) *EndpointConfig {
	e.Path = value
	return e
}

func (e *EndpointConfig) SetMethod(value string) *EndpointConfig {
	v, ok := curlx.MethodWithRequestBody[curlx.Method(value)]
	if !ok || !v {
		log.Panicf("Invalid method: %v", value)
	}
	e.Method = value
	return e
}

func (e *EndpointConfig) SetDescription(value string) *EndpointConfig {
	e.Description = value
	return e
}

func (e *EndpointConfig) SetQueryParams(value map[string]string) *EndpointConfig {
	e.QueryParams = value
	return e
}

func (e *EndpointConfig) AppendQueryParam(key string, value string) *EndpointConfig {
	if len(e.QueryParams) == 0 {
		e.SetQueryParams(make(map[string]string))
	}
	e.QueryParams[key] = value
	return e
}

func (e *EndpointConfig) SetPathParams(value map[string]string) *EndpointConfig {
	e.PathParams = value
	return e
}

func (e *EndpointConfig) AppendPathParam(key, value string) *EndpointConfig {
	if len(e.PathParams) == 0 {
		e.SetPathParams(make(map[string]string))
	}
	e.PathParams[key] = value
	return e
}

func (e *EndpointConfig) SetHeaders(value map[string]string) *EndpointConfig {
	e.Headers = value
	return e
}

func (e *EndpointConfig) AppendHeader(key, value string) *EndpointConfig {
	if len(e.Headers) == 0 {
		e.SetHeaders(make(map[string]string))
	}
	e.Headers[key] = value
	return e
}

func (e *EndpointConfig) SetBody(value map[string]interface{}) *EndpointConfig {
	e.Body = value
	return e
}

func (e *EndpointConfig) AppendBody(key, value string) *EndpointConfig {
	if len(e.Body) == 0 {
		e.SetBody(make(map[string]interface{}))
	}
	e.Body[key] = value
	return e
}

func (e *EndpointConfig) AppendBodyWith(key string, value interface{}) *EndpointConfig {
	if len(e.Body) == 0 {
		e.SetBody(make(map[string]interface{}))
	}
	e.Body[key] = value
	return e
}

func (e *EndpointConfig) SetRetry(value RetryConfig) *EndpointConfig {
	e.Retry = value
	return e
}

func (e *EndpointConfig) SetAuthentication(value AuthenticationConfig) *EndpointConfig {
	e.Authentication = value
	return e
}

func (e *EndpointConfig) SetTelegram(value telegram.TelegramConfig) *EndpointConfig {
	e.Telegram = value
	return e
}

func (e *EndpointConfig) Json() string {
	return utils.ToJson(e)
}

func EndpointValidator(e *EndpointConfig) {
	e.SetBaseURL(e.BaseURL).
		SetMethod(e.Method)
}

func GetEndpointSample() *EndpointConfig {
	e := NewEndpoint().
		SetEnabled(true).
		SetDebugMode(true).
		SetBaseURL("http://127.0.0.1:8080").
		SetTimeout(10*time.Second).
		SetPath("/api/v1/users").
		SetMethod("POST").
		SetDescription("Create new user").
		AppendHeader("Content-Type", "application/json").
		AppendBody("username", "tester").
		AppendBody("email", "tester@gmail.com").
		SetAuthentication(*GetAuthenticationSample()).
		SetRetry(*GetRetrySample()).
		SetTelegram(*telegram.GetTelegramConfigSample())
	return e
}

func NewApiRequest() *ApiRequestConfig {
	a := &ApiRequestConfig{}
	return a
}

func (a *ApiRequestConfig) SetBaseURL(value string) *ApiRequestConfig {
	u, err := url.Parse(value)
	if err != nil {
		log.Panicf("Invalid base_url: %v", err.Error())
	}
	a.BaseURL = u.String()
	return a
}

func (a *ApiRequestConfig) SetAuthentication(value AuthenticationConfig) *ApiRequestConfig {
	a.Authentication = value
	return a
}

func (a *ApiRequestConfig) SetHeaders(value map[string]string) *ApiRequestConfig {
	a.Headers = value
	return a
}

func (a *ApiRequestConfig) AppendHeader(key, value string) *ApiRequestConfig {
	if len(a.Headers) == 0 {
		a.SetHeaders(make(map[string]string))
	}
	a.Headers[key] = value
	return a
}

func (a *ApiRequestConfig) SetRetry(value RetryConfig) *ApiRequestConfig {
	a.Retry = value
	return a
}

func (a *ApiRequestConfig) SetEndpoints(value map[string]EndpointConfig) *ApiRequestConfig {
	a.Endpoints = value
	return a
}

func (a *ApiRequestConfig) AppendEndpoint(key string, endpoint EndpointConfig) *ApiRequestConfig {
	if len(a.Endpoints) == 0 {
		a.SetEndpoints(make(map[string]EndpointConfig))
	}
	a.Endpoints[key] = endpoint
	return a
}

func (a *ApiRequestConfig) SetTelegram(value telegram.TelegramConfig) *ApiRequestConfig {
	a.Telegram = value
	return a
}

func (a *ApiRequestConfig) SetKey(value string) *ApiRequestConfig {
	a.Key = value
	return a
}

func (a *ApiRequestConfig) Json() string {
	return utils.ToJson(a)
}

func GetApiRequestSample() *ApiRequestConfig {
	a := NewApiRequest().
		SetBaseURL("http://127.0.0.1:8080").
		SetAuthentication(*GetAuthenticationSample()).
		AppendHeader("Content-Type", "application/json").
		AppendEndpoint("a_endpoint", *GetEndpointSample()).
		AppendEndpoint("b_endpoint", *GetEndpointSample()).
		SetRetry(*GetRetrySample()).
		SetTelegram(*telegram.GetTelegramConfigSample())
	return a
}

func (a *ApiRequestConfig) AvailableEndpoint() bool {
	return len(a.Endpoints) > 0
}

func (a *ApiRequestConfig) AvailableHeader() bool {
	return len(a.Headers) > 0
}

func (a *ApiRequestConfig) GetEndpoint(key string) (EndpointConfig, error) {
	v, ok := a.Endpoints[key]
	if !ok {
		return EndpointConfig{}, fmt.Errorf("Endpoint not found")
	}
	return v, nil
}

func (e *EndpointConfig) Url() (string, error) {
	return e.UrlWith(e.BaseURL)
}

func (e *EndpointConfig) AvailableBody() bool {
	return len(e.Body) > 0
}

func (e *EndpointConfig) AvailableQueryParams() bool {
	return len(e.QueryParams) > 0
}

func (e *EndpointConfig) AvailablePathParams() bool {
	return len(e.PathParams) > 0
}

func (e *EndpointConfig) AvailableHeaders() bool {
	return len(e.Headers) > 0
}

func (e *EndpointConfig) UrlWith(baseURL string) (string, error) {
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

func (e *EndpointConfig) AvailableTimeout() bool {
	return isDuration(e.Timeout)
}

func (r *RetryConfig) AvailableRetryOnStatus() bool {
	return len(r.RetryOnStatus) > 0
}

func (a *ApiRequestConfig) FilterActivateEndpoints() *ApiRequestConfig {
	endpoints := make(map[string]EndpointConfig)
	for k, v := range a.Endpoints {
		if !v.IsEnabled {
			continue
		}
		endpoints[k] = v
	}
	a.SetEndpoints(endpoints)
	return a
}

func (a *ApiRequestConfig) CombineHeaders(e EndpointConfig) map[string]string {
	if a.AvailableHeader() {
		return coltx.MergeMapsString(a.Headers, e.Headers)
	}
	return e.Headers
}

func (a *ApiRequestConfig) CombineAuthentication(e EndpointConfig) AuthenticationConfig {
	if e.Authentication.IsEnabled {
		return e.Authentication
	}
	return a.Authentication
}

func (a *ApiRequestConfig) CombineRetry(e EndpointConfig) RetryConfig {
	if e.Retry.IsEnabled {
		return e.Retry
	}
	return a.Retry
}

func (a *ApiRequestConfig) CombineHostURL(e EndpointConfig) string {
	if utils.IsNotEmpty(e.BaseURL) {
		return e.BaseURL
	}
	return a.BaseURL
}

func (a *ApiRequestConfig) CombineUrl(e EndpointConfig) (string, error) {
	u, err := e.Url()
	if utils.IsEmpty(a.BaseURL) || utils.IsNotEmpty(e.BaseURL) {
		return u, err
	}
	return e.UrlWith(a.BaseURL)
}

func (a *ApiRequestConfig) CombineTelegram(e EndpointConfig) telegram.TelegramConfig {
	if e.Telegram.IsEnabled {
		return e.Telegram
	}
	return a.Telegram
}

func isDuration(t time.Duration) bool {
	return t != 0 && t > 0
}

func Get(node []ApiRequestConfig, key string) (ApiRequestConfig, bool) {
	if len(node) == 0 || utils.IsEmpty(key) {
		return ApiRequestConfig{}, false
	}
	var value ApiRequestConfig
	for _, v := range node {
		if v.Key == key {
			value = v
			break
		}
	}
	if utils.IsEmpty(value.Key) {
		return ApiRequestConfig{}, false
	}
	return value, true
}
