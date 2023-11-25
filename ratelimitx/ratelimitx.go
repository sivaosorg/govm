package ratelimitx

import (
	"fmt"
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewRateLimitConfig() *RateLimitConfig {
	return &RateLimitConfig{}
}

func (r *RateLimitConfig) SetEnabled(value bool) *RateLimitConfig {
	r.IsEnabled = value
	return r
}

func (r *RateLimitConfig) SetRate(value int) *RateLimitConfig {
	if value <= 0 {
		log.Panicf("Invalid rate: %v", value)
	}
	r.Rate = value
	return r
}

func (r *RateLimitConfig) SetMaxBurst(value int) *RateLimitConfig {
	if value <= 0 {
		log.Panicf("Invalid max_burst: %v", value)
	}
	r.MaxBurst = value
	return r
}

func (r *RateLimitConfig) Json() string {
	return utils.ToJson(r)
}

func GetRateLimitConfigSample() *RateLimitConfig {
	r := NewRateLimitConfig().SetEnabled(false).SetRate(100).SetMaxBurst(10)
	return r
}

func NewRateLimitOptionConfig() *rateLimitOptionConfig {
	return &rateLimitOptionConfig{
		MaxRetries: 2,
	}
}

func (r *rateLimitOptionConfig) SetMaxRetries(value int) *rateLimitOptionConfig {
	if value <= 0 {
		log.Panicf("Invalid max_retries: %v", value)
	}
	r.MaxRetries = value
	return r
}

func (r *rateLimitOptionConfig) Json() string {
	return utils.ToJson(r)
}

func NewMultiTenantRateLimitConfig() *MultiTenantRateLimitConfig {
	return &MultiTenantRateLimitConfig{}
}

func (m *MultiTenantRateLimitConfig) SetKey(value string) *MultiTenantRateLimitConfig {
	m.Key = value
	return m
}

func (m *MultiTenantRateLimitConfig) SetUsableDefault(value bool) *MultiTenantRateLimitConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantRateLimitConfig) SetConfig(value RateLimitConfig) *MultiTenantRateLimitConfig {
	m.Config = value
	return m
}

func (m *MultiTenantRateLimitConfig) SetConfigCursor(value *RateLimitConfig) *MultiTenantRateLimitConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantRateLimitConfig) SetOption(value rateLimitOptionConfig) *MultiTenantRateLimitConfig {
	m.Option = value
	return m
}

func (m *MultiTenantRateLimitConfig) Json() string {
	return utils.ToJson(m)
}

func GetMultiTenantRateLimitConfigSample() *MultiTenantRateLimitConfig {
	m := NewMultiTenantRateLimitConfig().
		SetKey("tenant_1").
		SetOption(*NewRateLimitOptionConfig()).
		SetConfigCursor(GetRateLimitConfigSample())
	return m
}

func NewClusterMultiTenantRateLimitConfig() *ClusterMultiTenantRateLimitConfig {
	return &ClusterMultiTenantRateLimitConfig{}
}

func (c *ClusterMultiTenantRateLimitConfig) SetClusters(values []MultiTenantRateLimitConfig) *ClusterMultiTenantRateLimitConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantRateLimitConfig) AppendClusters(values ...MultiTenantRateLimitConfig) *ClusterMultiTenantRateLimitConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantRateLimitConfig) Json() string {
	return utils.ToJson(c.Clusters)
}

func GetClusterMultiTenantRateLimitConfigSample() *ClusterMultiTenantRateLimitConfig {
	c := NewClusterMultiTenantRateLimitConfig().
		AppendClusters(*GetMultiTenantRateLimitConfigSample(), *GetMultiTenantRateLimitConfigSample().SetKey("tenant_2"))
	return c
}

func (c *ClusterMultiTenantRateLimitConfig) FindClusterBy(key string) (MultiTenantRateLimitConfig, error) {
	if utils.IsEmpty(key) {
		return *NewMultiTenantRateLimitConfig(), fmt.Errorf("Key is required")
	}
	if len(c.Clusters) == 0 {
		return *NewMultiTenantRateLimitConfig(), fmt.Errorf("No ratelimit cluster")
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantRateLimitConfig(), fmt.Errorf("The ratelimit cluster not found")
}
