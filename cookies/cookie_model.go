package cookies

import "time"

type CookieConfig struct {
	IsEnabled bool          `json:"enabled" yaml:"enabled"`
	Name      string        `json:"name" yaml:"name"`
	Value     string        `json:"value" yaml:"value"`
	Path      string        `json:"path" yaml:"path"`
	Domain    string        `json:"domain" yaml:"domain"`
	MaxAge    int           `json:"max_age" yaml:"max_age"`
	Secure    bool          `json:"secure" yaml:"secure"`
	HttpOnly  bool          `json:"http_only" yaml:"http_only"`
	Timeout   time.Duration `json:"timeout" yaml:"timeout"`
}

type cookieOptionConfig struct {
	MaxRetries int `json:"max_retries" yaml:"max_retries"`
}

type MultiTenantCookieConfig struct {
	Key             string             `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool               `json:"usable_default" yaml:"usable_default"`
	Config          CookieConfig       `json:"config" yaml:"config"`
	Option          cookieOptionConfig `json:"option" binding:"required" yaml:"option"`
}

type ClusterMultiTenantCookieConfig struct {
	Clusters []MultiTenantCookieConfig `json:"clusters,omitempty" yaml:"clusters"`
}
