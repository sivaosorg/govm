package ratelimitx

type RateLimitConfig struct {
	IsEnabled bool `json:"enabled" yaml:"enabled"`
	Rate      int  `json:"rate" yaml:"rate"`
	MaxBurst  int  `json:"max_burst" yaml:"max_burst"`
}

type rateLimitOptionConfig struct {
	MaxRetries int `json:"max_retries" yaml:"max_retries"`
}

type MultiTenantRateLimitConfig struct {
	Key             string                `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool                  `json:"usable_default" yaml:"usable_default"`
	Config          RateLimitConfig       `json:"config" yaml:"config"`
	Option          rateLimitOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantRateLimitConfig struct {
	Clusters []MultiTenantRateLimitConfig `json:"clusters,omitempty" yaml:"clusters"`
}
