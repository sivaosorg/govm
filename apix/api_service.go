package apix

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/sivaosorg/govm/blueprint"
	"github.com/sivaosorg/govm/bot/telegram"
	"github.com/sivaosorg/govm/coltx"
	"github.com/sivaosorg/govm/logger"
	"github.com/sivaosorg/govm/restify"
	"github.com/sivaosorg/govm/timex"
	"github.com/sivaosorg/govm/utils"
)

type ApiService interface {
	Do(client *restify.Client, endpoint EndpointConfig) (*restify.Response, error)
	DoAsyncWait(client *restify.Client, endpoints ...EndpointConfig)
	DoAsyncWaitWith(client *restify.Client, endpoints map[string]EndpointConfig)
	DoAsyncNoneWait(client *restify.Client, endpoints ...EndpointConfig)
	DoAsyncNoneWaitWith(client *restify.Client, endpoints map[string]EndpointConfig)
}

type apiServiceImpl struct {
	conf        ApiRequestConfig
	telegramSvc telegram.TelegramService
}

func NewApiService(conf ApiRequestConfig) ApiService {
	s := &apiServiceImpl{
		conf: conf,
	}
	return s
}

func (s *apiServiceImpl) Do(client *restify.Client, endpoint EndpointConfig) (*restify.Response, error) {
	return s.execute(client, endpoint)
}

func (s *apiServiceImpl) DoAsyncWait(client *restify.Client, endpoints ...EndpointConfig) {
	if len(endpoints) == 0 {
		return
	}
	var wg sync.WaitGroup
	start := time.Now()
	for _, v := range endpoints {
		if !v.IsEnabled {
			continue
		}
		wg.Add(1)
		go func(endpoint EndpointConfig) {
			defer wg.Done()
			_, err := s.Do(client, endpoint)
			if err == nil {
				// send notification
			}
		}(v)
	}
	wg.Wait()
	logger.Debugf(fmt.Sprintf("Size of endpoint(s): %v executed in %v", len(endpoints), time.Since(start)))
}

func (s *apiServiceImpl) DoAsyncWaitWith(client *restify.Client, endpoints map[string]EndpointConfig) {
	if len(endpoints) == 0 {
		return
	}
	var e []EndpointConfig
	for _, v := range endpoints {
		if !v.IsEnabled {
			continue
		}
		e = append(e, v)
	}
	s.DoAsyncWait(client, e...)
}

func (s *apiServiceImpl) DoAsyncNoneWait(client *restify.Client, endpoints ...EndpointConfig) {
	if len(endpoints) == 0 {
		return
	}
	start := time.Now()
	for _, v := range endpoints {
		if !v.IsEnabled {
			continue
		}
		go func(endpoint EndpointConfig) {
			_, err := s.Do(client, endpoint)
			if err == nil {
				// send notification
			}
		}(v)
	}
	logger.Debugf(fmt.Sprintf("Size of endpoint(s): %v executed in %v", len(endpoints), time.Since(start)))
}

func (s *apiServiceImpl) DoAsyncNoneWaitWith(client *restify.Client, endpoints map[string]EndpointConfig) {
	if len(endpoints) == 0 {
		return
	}
	var e []EndpointConfig
	for _, v := range endpoints {
		if !v.IsEnabled {
			continue
		}
		e = append(e, v)
	}
	s.DoAsyncNoneWait(client, e...)
}

func (s *apiServiceImpl) execute(client *restify.Client, endpoint EndpointConfig) (*restify.Response, error) {
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
	telegramConf := s.conf.CombineTelegram(endpoint)
	options := telegram.NewTelegramOptionConfig().SetType(telegram.ModeMarkdown)
	telegramSvc := telegram.NewTelegramService(telegramConf, *options)
	s.telegramSvc = telegramSvc

	client.SetBaseURL(host)
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
					go s.alert(endpoint, _response, err)
				}
				return confirm
			} else {
				confirm := (_response.StatusCode() >= http.StatusBadRequest && _response.StatusCode() <= http.StatusNetworkAuthenticationRequired) || err != nil
				if confirm {
					// send notification
					// use goroutine
					go s.alert(endpoint, _response, err)
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
	if endpoint.AvailableBody() {
		request.SetBody(endpoint.Body)
	}
	response, err = request.Execute(endpoint.Method, fullURL)
	if endpoint.TelegramOptions.IsEnabledPingResponse {
		go s.alert(endpoint, response, err)
	}
	return response, err
}

func (s *apiServiceImpl) alert(endpoint EndpointConfig, response *restify.Response, err error) {
	var message strings.Builder
	if err != nil || response.IsError() {
		icon, _ := blueprint.TypeIcons[blueprint.TypeError]
		message.WriteString(fmt.Sprintf("%v ", icon))
	}
	if err == nil && response.IsSuccess() {
		icon, _ := blueprint.TypeIcons[blueprint.TypeNotification]
		message.WriteString(fmt.Sprintf("%v ", icon))
	}
	message.WriteString("API REST HTTP\n")
	message.WriteString(fmt.Sprintf("tz: `%s` (received at: `%v`)\n\n",
		time.Now().Format(timex.TimeFormat20060102150405), response.ReceivedAt()))
	if utils.IsNotEmpty(endpoint.Description) {
		message.WriteString(fmt.Sprintf("decs: `%s`\n", endpoint.Description))
	}
	message.WriteString(fmt.Sprintf("(`%s`) url: `%s`\n", response.Request.Method, response.Request.URL))
	message.WriteString("\n---\n")
	if endpoint.AvailableHeaders() && !endpoint.TelegramOptions.SkipMessageHeader {
		message.WriteString(fmt.Sprintf("header(s): \n\t`%s`\n", coltx.MapString2Table(endpoint.Headers)))
	}
	if endpoint.AvailableQueryParams() && !endpoint.TelegramOptions.SkipMessageQueryParam {
		message.WriteString(fmt.Sprintf("query param(s): `%s`\n", coltx.MapString2Table(endpoint.QueryParams)))
	}
	if endpoint.AvailablePathParams() && !endpoint.TelegramOptions.SkipMessagePathParam {
		message.WriteString(fmt.Sprintf("path param(s): `%s`\n", coltx.MapString2Table(endpoint.PathParams)))
	}
	if endpoint.AvailableBody() && !endpoint.TelegramOptions.SkipMessageRequestBody {
		message.WriteString(fmt.Sprintf("request body: `%s`\n", utils.ToJson(endpoint.Body)))
	}
	if endpoint.Retry.AvailableRetryOnStatus() {
		message.WriteString(fmt.Sprintf("retry on status: `%s`\n", utils.ToJson(endpoint.Retry.RetryOnStatus)))
	}
	message.WriteString("\n---\n")
	message.WriteString(fmt.Sprintf("status code: %v\n", response.StatusCode()))
	if utils.IsNotEmpty(response.String()) && !endpoint.TelegramOptions.SkipMessageResponseBody {
		message.WriteString(fmt.Sprintf("response: `%v`\n", response.String()))
	}
	message.WriteString(fmt.Sprintf("no. attempt: %v\n", response.Request.Attempt))
	message.WriteString(fmt.Sprintf("duration: `%v`\n", response.Time()))
	if err != nil {
		message.WriteString(fmt.Sprintf("error(r): `%v`\n", err.Error()))
	}
	s.telegramSvc.SendMessage(message.String())
}
