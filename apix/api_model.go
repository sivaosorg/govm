package apix

import "time"

type Authentication struct {
	IsEnabled bool   `json:"enabled" yaml:"enabled"`
	Type      string `json:"type,omitempty" yaml:"type"`
	Token     string `json:"-" yaml:"token"`
	Username  string `json:"username,omitempty" yaml:"username"`
	Password  string `json:"-" yaml:"password"`
}

type Retry struct {
	IsEnabled       bool          `json:"enabled" yaml:"enabled"`
	MaxAttempts     int           `json:"max_attempts" yaml:"max_attempts"`
	InitialInterval time.Duration `json:"initial_interval" yaml:"initial_interval"`
	MaxInterval     time.Duration `json:"max_interval" yaml:"max_interval"`
	BackoffFactor   int           `json:"backoff_factor" yaml:"backoff_factor"`
	RetryOnStatus   []int         `json:"retry_on_status" yaml:"retry_on_status"`
}

type Endpoint struct {
	IsEnabled      bool                   `json:"enabled" yaml:"enabled"`
	DebugMode      bool                   `json:"debug_mode" yaml:"debug_mode"`
	BaseURL        string                 `json:"base_url" yaml:"base_url"`
	Timeout        time.Duration          `json:"timeout" yaml:"timeout"`
	Path           string                 `json:"path" yaml:"path"`
	Method         string                 `json:"method" yaml:"method"`
	Description    string                 `json:"description" yaml:"description"`
	QueryParams    map[string]string      `json:"query_params" yaml:"query_params"`
	PathParams     map[string]string      `json:"path_params" yaml:"path_params"`
	Headers        map[string]string      `json:"headers" yaml:"headers"`
	Body           map[string]interface{} `json:"body" yaml:"body"`
	Retry          Retry                  `json:"retry" yaml:"retry"`
	Authentication Authentication         `json:"authentication" yaml:"authentication"`
	Repeat         int                    `json:"repeat" yaml:"repeat"`
}

type ApiRequest struct {
	BaseURL        string              `json:"base_url" yaml:"base_url"`
	Authentication Authentication      `json:"authentication" yaml:"authentication"`
	Headers        map[string]string   `json:"headers" yaml:"headers"`
	Endpoints      map[string]Endpoint `json:"endpoints" yaml:"endpoints"`
	Retry          Retry               `json:"retry" yaml:"retry"`
}
