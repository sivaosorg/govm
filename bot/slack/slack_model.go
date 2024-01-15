package slack

import "time"

type SlackConfig struct {
	IsEnabled bool          `json:"enabled" yaml:"enabled"`
	DebugMode bool          `json:"debug_mode" yaml:"debug_mode"`
	ChannelId []string      `json:"channel_id" binding:"required" yaml:"channel_id"`
	Token     string        `json:"-" binding:"required" yaml:"token"`
	Timeout   time.Duration `json:"timeout" yaml:"timeout"`
}

type slackOptionConfig struct {
	MaxRetries int `json:"max_retries" yaml:"max-retries"`
}

type MultiTenantSlackConfig struct {
	Key             string            `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool              `json:"usable_default" yaml:"usable_default"`
	Config          SlackConfig       `json:"config" yaml:"config"`
	Option          slackOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantSlackConfig struct {
	Clusters []MultiTenantSlackConfig `json:"clusters,omitempty" yaml:"clusters"`
}
