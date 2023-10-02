package redisx

import "time"

type RedisConfig struct {
	IsEnabled bool          `json:"enabled" yaml:"enabled"`
	DebugMode bool          `json:"debug_mode" yaml:"debug-mode"`
	UrlConn   string        `json:"url_conn" binding:"required" yaml:"url-conn"`
	Password  string        `json:"-" binding:"required" yaml:"password"`
	Database  string        `json:"database" binding:"required" yaml:"database"`
	Timeout   time.Duration `json:"-" yaml:"-"`
}

type redisOptionConfig struct {
}

type MultiTenantRedisConfig struct {
	Key             string            `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool              `json:"usable_default" yaml:"usable_default"`
	Config          RedisConfig       `json:"config" yaml:"config"`
	Option          redisOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantRedisConfig struct {
	Clusters []MultiTenantRedisConfig `json:"clusters,omitempty" yaml:"clusters"`
}
