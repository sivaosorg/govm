package curlx

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewCurlxContext() *CurlxContext {
	c := &CurlxContext{}
	c.SetMaxRetries(2)
	c.SetTimeout(15 * time.Second)
	c.SetMaxIdleConn(10)
	c.SetMaxIdleConnPerRequest(3)
	c.SetRetryContext(make([]RetryContext, 0))
	return c
}

func NewCurlxRequest() *CurlxRequest {
	c := &CurlxRequest{}
	c.SetDebugMode(true)
	c.SetQueryParams(make(map[string]string))
	c.SetHeaders(make(map[string]string))
	c.SetResponseBody(new(interface{}))
	c.SetResponseError(new(interface{}))
	c.SetCookies(make([]*http.Cookie, 0))
	c.SetAttachFileField(FieldNameFileForm)
	return c
}

func (c *CurlxContext) SetBaseURL(value string) *CurlxContext {
	if utils.IsEmpty(value) {
		log.Panicf("Base URL is required")
	}
	c.BaseURL = value
	return c
}

func (c *CurlxContext) SetMaxRetries(value int) *CurlxContext {
	if value <= 0 {
		log.Panicf("Invalid max-retries: %v", value)
	}
	c.MaxRetries = value
	return c
}

func (c *CurlxContext) SetMaxIdleConn(value int) *CurlxContext {
	if value <= 0 {
		log.Panicf("Invalid max-idle-conn: %v", value)
	}
	c.MaxIdleConns = value
	return c
}

func (c *CurlxContext) SetMaxIdleConnPerRequest(value int) *CurlxContext {
	if value <= 0 {
		log.Panicf("Invalid max-idle-conn-per-request: %v", value)
	}
	c.MaxIdleConnsPerRequest = value
	return c
}

func (c *CurlxContext) SetTimeout(value time.Duration) *CurlxContext {
	c.Timeout = value
	return c
}

func (c *CurlxContext) SetRetryContext(value []RetryContext) *CurlxContext {
	c.RetryContext = value
	return c
}

func (c *CurlxContext) AppendRetryContext(value ...RetryContext) *CurlxContext {
	c.RetryContext = append(c.RetryContext, value...)
	return c
}

func (c *CurlxContext) Json() string {
	return utils.ToJson(c)
}

func CurlxContextValidator(c *CurlxContext) {
	c.SetBaseURL(c.BaseURL).
		SetMaxIdleConn(c.MaxIdleConns).
		SetMaxIdleConnPerRequest(c.MaxIdleConnsPerRequest).
		SetMaxRetries(c.MaxRetries)
}

func (c *CurlxRequest) SetMethod(value Method) *CurlxRequest {
	if utils.IsEmpty(string(value)) {
		log.Panicf("Invalid method")
	}
	c.Method = value
	return c
}

func (c *CurlxRequest) SetEndpoint(value string) *CurlxRequest {
	c.Endpoint = value
	return c
}

func (c *CurlxRequest) SetAttachment(value string) *CurlxRequest {
	if utils.IsEmpty(value) {
		return c
	}
	c.Attachment = filepath.Clean(value)
	return c
}

func (c *CurlxRequest) SetRequestBody(value interface{}) *CurlxRequest {
	c.RequestBody = value
	return c
}

func (c *CurlxRequest) SetResponseBody(value interface{}) *CurlxRequest {
	c.ResponseBody = value
	return c
}

func (c *CurlxRequest) SetQueryParams(value map[string]string) *CurlxRequest {
	c.QueryParams = value
	return c
}

func (c *CurlxRequest) AppendQueryParam(key string, value string) *CurlxRequest {
	c.QueryParams[key] = value
	return c
}

func (c *CurlxRequest) AppendQueryParamWith(key string, value interface{}) *CurlxRequest {
	if utils.IsPrimitiveType(value) {
		c.QueryParams[key] = fmt.Sprintf("%v", value)
	} else {
		c.QueryParams[key] = utils.ToJson(value)
	}
	return c
}

func (c *CurlxRequest) SetQueryParamsWith(value map[string]interface{}) *CurlxRequest {
	if len(value) == 0 {
		return c
	}
	for k, v := range value {
		c.AppendQueryParamWith(k, v)
	}
	return c
}

func (c *CurlxRequest) SetHeaders(value map[string]string) *CurlxRequest {
	c.Headers = value
	return c
}

func (c *CurlxRequest) AppendHeader(key string, value string) *CurlxRequest {
	c.Headers[key] = value
	return c
}

func (c *CurlxRequest) AppendHeaderWith(key string, value interface{}) *CurlxRequest {
	if utils.IsPrimitiveType(value) {
		c.Headers[key] = fmt.Sprintf("%v", value)
	} else {
		c.Headers[key] = utils.ToJson(value)
	}
	return c
}

func (c *CurlxRequest) SetCookies(value []*http.Cookie) *CurlxRequest {
	c.Cookies = value
	return c
}

func (c *CurlxRequest) AppendCookie(values ...*http.Cookie) *CurlxRequest {
	c.Cookies = append(c.Cookies, values...)
	return c
}

func (c *CurlxRequest) AppendCookieKV(key string, value interface{}) *CurlxRequest {
	var v string
	if utils.IsPrimitiveType(value) {
		v = fmt.Sprintf("%v", value)
	} else {
		v = utils.ToJson(value)
	}
	c.Cookies = append(c.Cookies, &http.Cookie{
		Name:  key,
		Value: v,
	})
	return c
}

func (c *CurlxRequest) SetDebugMode(value bool) *CurlxRequest {
	c.DebugMode = value
	return c
}

func (c *CurlxRequest) SetResponseError(value interface{}) *CurlxRequest {
	c.ResponseError = value
	return c
}

func (c *CurlxRequest) SetAttachFileField(value string) *CurlxRequest {
	if utils.IsEmpty(value) {
		log.Panicf("Invalid attach_file_field: %v", value)
	}
	c.AttachFileField = value
	return c
}

func (c *CurlxRequest) Json() string {
	return utils.ToJson(c)
}

func CurlxRequestValidator(c *CurlxRequest) {
	c.SetMethod(c.Method)
}
