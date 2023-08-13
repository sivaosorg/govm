package curlx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/sivaosorg/govm/entity"
	"github.com/sivaosorg/govm/logger"
	"github.com/sivaosorg/govm/utils"
)

type CurlxService interface {
	retry(attempt int, response *http.Response, err error) bool
	doRequest(req *http.Request, lastError *error) (*http.Response, error)
	getUrlCompleted() string
	Fetch() error
	FetchWithProgress(callback ProgressCallback) error
	ProgressCallbackFunc() ProgressCallback
}

type curlxServiceImpl struct {
	context CurlxContext `json:"-"`
	request CurlxRequest `json:"-"`
}

func NewCurlxService(context CurlxContext, request CurlxRequest) CurlxService {
	s := &curlxServiceImpl{
		context: context,
		request: request,
	}
	return s
}

func (s *curlxServiceImpl) retry(attempt int, response *http.Response, err error) bool {
	if s.request.DebugMode {
		message := fmt.Sprintf("[CURL]. Request %v attempt retrying... <%v> time(s), total_remaining_attempt %v", s.getUrlCompleted(), attempt, s.context.MaxRetries-attempt)
		logger.Debugf(message)
		message = fmt.Sprintf("[CURL]. Request %v attempt retry configured %v time(s)", s.getUrlCompleted(), s.context.MaxRetries)
		logger.Debugf(message)
	}
	if attempt >= s.context.MaxRetries {
		return false
	}
	if response != nil && entity.IsStatusCodeSuccess(response.StatusCode) {
		return false
	}
	if len(s.context.RetryContext) == 0 {
		return false
	}
	for _, condition := range s.context.RetryContext {
		attempt++
		if condition(attempt, response, err) {
			return true
		}
	}
	return false
}

func (s *curlxServiceImpl) doRequest(request *http.Request, _err *error) (*http.Response, error) {
	CurlxContextValidator(&s.context)
	CurlxRequestValidator(&s.request)
	if s.request.ResponseBody == nil {
		s.request.SetResponseBody(new(interface{}))
	}
	if s.request.ResponseError == nil {
		s.request.SetResponseError(new(interface{}))
	}
	if s.context.MaxIdleConns <= 0 {
		s.context.SetMaxIdleConn(10)
	}
	if s.context.MaxIdleConnsPerRequest <= 0 {
		s.context.SetMaxIdleConnPerRequest(3)
	}
	transport := &http.Transport{
		MaxIdleConns:        s.context.MaxIdleConns,
		MaxIdleConnsPerHost: s.context.MaxIdleConnsPerRequest,
	}
	client := http.Client{
		Transport: transport,
		Timeout:   s.context.Timeout,
	}
	raw, err := client.Do(request)
	if err != nil {
		*_err = err
		if raw != nil {
			bytes, _ := ioutil.ReadAll(raw.Body)
			s.request.SetResponseError(string(bytes))
		}
		return raw, err
	}
	defer raw.Body.Close()
	if s.request.DebugMode {
		v := fmt.Sprintf("[CURL]. Request %v response status code %v", s.getUrlCompleted(), raw.StatusCode)
		logger.Debugf(v)
	}
	if entity.IsStatusCodeFailure(raw.StatusCode) {
		decoder := json.NewDecoder(raw.Body)
		_ = decoder.Decode(s.request.ResponseError)
		*_err = fmt.Errorf("(Conn alive) request %v got failed with response status code %d", s.getUrlCompleted(), raw.StatusCode)
		return raw, *_err
	}
	decoder := json.NewDecoder(raw.Body)
	err = decoder.Decode(s.request.ResponseBody)
	if err != nil {
		*_err = err
		return raw, err
	}
	return raw, nil
}

func (s *curlxServiceImpl) getUrlCompleted() string {
	_url := fmt.Sprintf("%s%s", s.context.BaseURL, s.request.Endpoint)
	if len(s.request.QueryParams) > 0 {
		queryParams := url.Values{}
		for key, value := range s.request.QueryParams {
			queryParams.Add(key, value)
		}
		_url += "?" + queryParams.Encode()
	}
	return _url
}

func (s *curlxServiceImpl) Fetch() error {
	CurlxContextValidator(&s.context)
	CurlxRequestValidator(&s.request)
	url := fmt.Sprintf("%s%s", s.context.BaseURL, s.request.Endpoint)
	if s.request.DebugMode {
		message := fmt.Sprintf("[CURL]. Request %v sending...", s.getUrlCompleted())
		logger.Debugf(message)
	}
	_buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(_buffer)
	// adding headers
	if len(s.request.Headers) > 0 {
		for key, value := range s.request.Headers {
			writer.WriteField(key, value)
			if s.request.DebugMode {
				message := fmt.Sprintf("[CURL]. Request %v added header key [%v = %v]", s.getUrlCompleted(), key, value)
				logger.Debugf(message)
			}
		}
	}
	// adding attachment
	if utils.IsNotEmpty(s.request.Attachment) {
		file, err := os.Open(s.request.Attachment)
		if err != nil {
			return err
		}
		defer file.Close()
		part, err := writer.CreateFormFile("file", s.request.Attachment)
		if err != nil {
			return err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return err
		}
		if s.request.DebugMode {
			v := fmt.Sprintf("[CURL]. Request %v added attachment filename:: %v", s.getUrlCompleted(), file.Name())
			logger.Debugf(v)
		}
	}
	writer.Close()
	// adding header content-type
	_request, err := http.NewRequest(string(s.request.Method), url, _buffer)
	if err != nil {
		return err
	}
	if utils.IsNotEmpty(s.request.Attachment) {
		_request.Header.Set("Content-Type", writer.FormDataContentType())
	}
	// adding cookies
	if len(s.request.Cookies) > 0 {
		for _, v := range s.request.Cookies {
			_request.AddCookie(v)
			if s.request.DebugMode {
				message := fmt.Sprintf("[CURL]. Request %v added cookie key [%v = %v]", s.getUrlCompleted(), v.Name, v.Value)
				logger.Debugf(message)
			}
		}
	}
	// adding query params
	if len(s.request.QueryParams) > 0 {
		queries := _request.URL.Query()
		for key, value := range s.request.QueryParams {
			queries.Add(key, value)
			if s.request.DebugMode {
				message := fmt.Sprintf("[CURL]. Request %v added query params key [%v = %v]", url, key, value)
				logger.Debugf(message)
			}
		}
		_request.URL.RawQuery = queries.Encode()
	}
	var _err error
	for attempt := 0; attempt <= s.context.MaxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(0 * time.Second)
		}
		_err = nil
		_response, err := s.doRequest(_request, &_err)
		if s.retry(attempt, _response, err) {
			continue
		}
		if err == nil {
			return nil
		}
	}
	return _err
}

func (s *curlxServiceImpl) FetchWithProgress(callback ProgressCallback) error {
	CurlxContextValidator(&s.context)
	CurlxRequestValidator(&s.request)
	url := fmt.Sprintf("%s%s", s.context.BaseURL, s.request.Endpoint)
	if s.request.DebugMode {
		message := fmt.Sprintf("[CURL]. Request %v sending...", s.getUrlCompleted())
		logger.Debugf(message)
	}
	if callback == nil {
		callback = s.ProgressCallbackFunc()
	}
	_buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(_buffer)
	// adding headers
	if len(s.request.Headers) > 0 {
		for key, value := range s.request.Headers {
			writer.WriteField(key, value)
			if s.request.DebugMode {
				message := fmt.Sprintf("[CURL]. Request %v added header key [%v = %v]", s.getUrlCompleted(), key, value)
				logger.Debugf(message)
			}
		}
	}
	// adding attachment
	if utils.IsNotEmpty(s.request.Attachment) {
		file, err := os.Open(s.request.Attachment)
		if err != nil {
			return err
		}
		defer file.Close()
		part, err := writer.CreateFormFile("file", s.request.Attachment)
		if err != nil {
			return err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return err
		}
		if s.request.DebugMode {
			v := fmt.Sprintf("[CURL]. Request %v added attachment filename:: %v", s.getUrlCompleted(), file.Name())
			logger.Debugf(v)
		}
	}
	writer.Close()
	// adding header content-type
	_request, err := http.NewRequest(string(s.request.Method), url, _buffer)
	if err != nil {
		return err
	}
	if utils.IsNotEmpty(s.request.Attachment) {
		_request.Header.Set("Content-Type", writer.FormDataContentType())
	}
	// adding cookies
	if len(s.request.Cookies) > 0 {
		for _, v := range s.request.Cookies {
			_request.AddCookie(v)
			if s.request.DebugMode {
				message := fmt.Sprintf("[CURL]. Request %v added cookie key [%v = %v]", s.getUrlCompleted(), v.Name, v.Value)
				logger.Debugf(message)
			}
		}
	}
	// adding query params
	if len(s.request.QueryParams) > 0 {
		queries := _request.URL.Query()
		for key, value := range s.request.QueryParams {
			queries.Add(key, value)
			if s.request.DebugMode {
				message := fmt.Sprintf("[CURL]. Request %v added query params key [%v = %v]", url, key, value)
				logger.Debugf(message)
			}
		}
		_request.URL.RawQuery = queries.Encode()
	}
	var _err error
	progress := 0.0
	totalAttempts := s.context.MaxRetries + 1
	avgDuration := time.Duration(0)
	now := time.Now()
	for attempt := 0; attempt <= s.context.MaxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(0 * time.Second)
		}
		start := time.Now()
		_err = nil
		_response, err := s.doRequest(_request, &_err)

		// calculation progress
		avgDuration += time.Since(start)
		avgDuration /= time.Duration(attempt + 1)
		estimatedTotalTime := avgDuration * time.Duration(totalAttempts)
		elapsed := time.Since(now)
		duration := estimatedTotalTime - elapsed
		if duration < 0 {
			duration = 0
		}
		progress = (float64(attempt) / float64(totalAttempts)) * 100.0
		if duration > 0 {
			_progress := (elapsed.Seconds() / estimatedTotalTime.Seconds()) * 100.0
			if _progress > progress {
				progress = _progress
			}
		}
		callback(attempt, progress, duration)
		if s.retry(attempt, _response, err) {
			continue
		}
		if err == nil {
			return nil
		}
	}
	return _err
}

func (s *curlxServiceImpl) ProgressCallbackFunc() ProgressCallback {
	return func(attempt int, percentage float64, duration time.Duration) {
		if s.request.DebugMode {
			message := fmt.Sprintf("[CURL]. Request %v progress: %.2f(%s) and remain duration: %v (attempt. no <%v>)", s.getUrlCompleted(), percentage, "%%", duration, attempt)
			logger.Warnf(message)
		}
	}
}
