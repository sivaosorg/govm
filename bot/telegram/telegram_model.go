package telegram

import "time"

type TelegramFormatType string

type TelegramConfig struct {
	IsEnabled bool          `json:"enabled" yaml:"enabled"`
	DebugMode bool          `json:"debug_mode" yaml:"debug_mode"`
	ChatID    []int64       `json:"chat_id" binding:"required" yaml:"chat_id"`
	Token     string        `json:"-" binding:"required" yaml:"token"`
	Timeout   time.Duration `json:"-" yaml:"-"`
}

type telegramOptionConfig struct {
	Type       TelegramFormatType `json:"type" binding:"required" yaml:"type"`
	MaxRetries int                `json:"max_retries" yaml:"max-retries"`
}

type MultiTenantTelegramConfig struct {
	Key             string               `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool                 `json:"usable_default" yaml:"usable_default"`
	Config          TelegramConfig       `json:"config" yaml:"config"`
	Option          telegramOptionConfig `json:"option" binding:"required" yaml:"option"`
}

type ClusterMultiTenantTelegramConfig struct {
	Clusters []MultiTenantTelegramConfig `json:"clusters,omitempty" yaml:"clusters"`
}

// Components

// Table of contents
// Button
type button struct {
	Text string `json:"text"`
	Url  string `json:"url,omitempty"`
}

// Inline Keyboard
type inlineKeyboard struct {
	Text    string `json:"text"`
	Url     string `json:"url,omitempty"`
	Payload string `json:"callback_data,omitempty"`
}
