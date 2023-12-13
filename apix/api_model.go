package apix

import (
	"time"

	"github.com/sivaosorg/govm/bot/telegram"
)

type AuthenticationConfig struct {
	IsEnabled bool   `json:"enabled" yaml:"enabled"`
	Type      string `json:"type,omitempty" yaml:"type"`
	Token     string `json:"-" yaml:"token"`
	Username  string `json:"username,omitempty" yaml:"username"`
	Password  string `json:"-" yaml:"password"`
}

type RetryConfig struct {
	IsEnabled       bool          `json:"enabled" yaml:"enabled"`
	MaxAttempts     int           `json:"max_attempts" yaml:"max_attempts"`
	InitialInterval time.Duration `json:"initial_interval" yaml:"initial_interval"`
	MaxInterval     time.Duration `json:"max_interval" yaml:"max_interval"`
	BackoffFactor   int           `json:"backoff_factor" yaml:"backoff_factor"`
	RetryOnStatus   []int         `json:"retry_on_status" yaml:"retry_on_status"`
}

type EndpointConfig struct {
	IsEnabled       bool                    `json:"enabled" yaml:"enabled"`
	DebugMode       bool                    `json:"debug_mode" yaml:"debug_mode"`
	BaseURL         string                  `json:"base_url" yaml:"base_url"`
	Timeout         time.Duration           `json:"timeout" yaml:"timeout"`
	Path            string                  `json:"path" yaml:"path"`
	Method          string                  `json:"method" yaml:"method"`
	Description     string                  `json:"description" yaml:"description"`
	QueryParams     map[string]string       `json:"query_params" yaml:"query_params"`
	PathParams      map[string]string       `json:"path_params" yaml:"path_params"`
	Headers         map[string]string       `json:"headers" yaml:"headers"`
	Body            map[string]interface{}  `json:"body" yaml:"body"`
	Retry           RetryConfig             `json:"retry" yaml:"retry"`
	Authentication  AuthenticationConfig    `json:"authentication" yaml:"authentication"`
	Telegram        telegram.TelegramConfig `json:"telegram" yaml:"telegram"`
	TelegramOptions EndpointOptionsConfig   `json:"telegram_options" yaml:"telegram_options"`
}

type EndpointOptionsConfig struct {
	IsEnabledPingResponse   bool `json:"enabled_ping_response" yaml:"enabled_ping_response"`
	SkipMessageHeader       bool `json:"skip_message_header" yaml:"skip_message_header"`
	SkipMessageRequestBody  bool `json:"skip_message_request_body" yaml:"skip_message_request_body"`
	SkipMessageResponseBody bool `json:"skip_message_response_body" yaml:"skip_message_response_body"`
	SkipMessageQueryParam   bool `json:"skip_message_query_param" yaml:"skip_message_query_param"`
	SkipMessagePathParam    bool `json:"skip_message_path_param" yaml:"skip_message_path_param"`
}

type ApiRequestConfig struct {
	BaseURL        string                    `json:"base_url" yaml:"base_url"`
	Key            string                    `json:"key" yaml:"key"`
	Authentication AuthenticationConfig      `json:"authentication" yaml:"authentication"`
	Headers        map[string]string         `json:"headers" yaml:"headers"`
	Endpoints      map[string]EndpointConfig `json:"endpoints" yaml:"endpoints"`
	Retry          RetryConfig               `json:"retry" yaml:"retry"`
	Telegram       telegram.TelegramConfig   `json:"telegram" yaml:"telegram"`
}
