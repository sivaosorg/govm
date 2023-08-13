package curlx

import (
	"net/http"
	"time"
)

type RetryContext func(int, *http.Response, error) bool

type CurlxContext struct {
	BaseURL                string         `json:"base_url" binding:"required" yaml:"base-url"`          // Set base-url, just like the host
	MaxRetries             int            `json:"max_retries" yaml:"max-retires"`                       // Set max-retries if the request got failure
	MaxIdleConns           int            `json:"max_idle_conns" yaml:"max-idle-conns"`                 // Maximum idle connections in the pool
	MaxIdleConnsPerRequest int            `json:"max_idle_conns_per_req" yaml:"max-idle-conns-per-req"` // Maximum idle connections per request
	Timeout                time.Duration  `json:"timeout" yaml:"timeout"`                               // Request timeout
	RetryContext           []RetryContext `json:"-"`                                                    // The condition to retry request
}

type CurlxRequest struct {
	DebugMode     bool              `json:"debug_mode" yaml:"debug-mode"`
	Method        Method            `json:"method" binding:"required" yaml:"method"`
	Endpoint      string            `json:"endpoint" yaml:"endpoint"`
	Attachment    string            `json:"attachment,omitempty" yaml:"attachment"`
	RequestBody   interface{}       `json:"request_body" yaml:"request-body"`
	ResponseBody  interface{}       `json:"response_body,omitempty" yaml:"response-body"`
	ResponseError interface{}       `json:"response_error,omitempty" yaml:"response-error"`
	QueryParams   map[string]string `json:"query_params" yaml:"query-params"`
	Headers       map[string]string `json:"headers" yaml:"headers"`
	Cookies       []*http.Cookie    `json:"cookies,omitempty"`
}
