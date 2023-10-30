package cookies

import (
	"fmt"
	"log"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewCookieConfig() *CookieConfig {
	return &CookieConfig{}
}

func (c *CookieConfig) SetName(value string) *CookieConfig {
	c.Name = value
	return c
}

func (c *CookieConfig) SetValue(value string) *CookieConfig {
	c.Value = value
	return c
}

func (c *CookieConfig) SetPath(value string) *CookieConfig {
	c.Path = value
	return c
}

func (c *CookieConfig) SetDomain(value string) *CookieConfig {
	c.Domain = value
	return c
}

func (c *CookieConfig) SetMaxAge(value int) *CookieConfig {
	if value < 0 {
		log.Panicf("Invalid max_age: %v", value)
	}
	c.MaxAge = value
	return c
}

func (c *CookieConfig) SetSecure(value bool) *CookieConfig {
	c.Secure = value
	return c
}

func (c *CookieConfig) SetHttpOnly(value bool) *CookieConfig {
	c.HttpOnly = value
	return c
}

func (c *CookieConfig) SetTimeout(value time.Duration) *CookieConfig {
	c.Timeout = value
	return c
}

func (c *CookieConfig) SetEnabled(value bool) *CookieConfig {
	c.IsEnabled = value
	return c
}

func (c *CookieConfig) Json() string {
	return utils.ToJson(c)
}

func GetCookieConfigSample() *CookieConfig {
	c := NewCookieConfig().
		SetEnabled(true).
		SetHttpOnly(true).
		SetSecure(false).
		SetName("user").
		SetMaxAge(86400).
		SetTimeout(10 * time.Second).
		SetPath("/")
	return c
}

func NewCookieOptionConfig() *cookieOptionConfig {
	return &cookieOptionConfig{
		MaxRetries: 2,
	}
}

func NewMultiTenantCookieConfig() *MultiTenantCookieConfig {
	return &MultiTenantCookieConfig{}
}

func (m *MultiTenantCookieConfig) SetKey(value string) *MultiTenantCookieConfig {
	m.Key = value
	return m
}

func (m *MultiTenantCookieConfig) SetUsableDefault(value bool) *MultiTenantCookieConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantCookieConfig) SetConfig(value CookieConfig) *MultiTenantCookieConfig {
	m.Config = value
	return m
}

func (m *MultiTenantCookieConfig) SetConfigCursor(value *CookieConfig) *MultiTenantCookieConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantCookieConfig) SetOption(value cookieOptionConfig) *MultiTenantCookieConfig {
	m.Option = value
	return m
}

func (m *MultiTenantCookieConfig) Json() string {
	return utils.ToJson(m)
}

func GetMultiTenantCookieConfigSample() *MultiTenantCookieConfig {
	m := NewMultiTenantCookieConfig().
		SetKey("tenant_1").
		SetUsableDefault(true).
		SetConfig(*GetCookieConfigSample()).
		SetOption(*NewCookieOptionConfig())
	return m
}

func NewClusterMultiTenantCookieConfig() *ClusterMultiTenantCookieConfig {
	return &ClusterMultiTenantCookieConfig{}
}

func (c *ClusterMultiTenantCookieConfig) SetClusters(value []MultiTenantCookieConfig) *ClusterMultiTenantCookieConfig {
	c.Clusters = value
	return c
}

func (c *ClusterMultiTenantCookieConfig) AppendClusters(values ...MultiTenantCookieConfig) *ClusterMultiTenantCookieConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantCookieConfig) Json() string {
	return utils.ToJson(c.Clusters)
}

func (c *ClusterMultiTenantCookieConfig) FindClusterBy(key string) (MultiTenantCookieConfig, error) {
	if utils.IsEmpty(key) {
		return *NewMultiTenantCookieConfig(), fmt.Errorf("Key is required")
	}
	if len(c.Clusters) == 0 {
		return *NewMultiTenantCookieConfig(), fmt.Errorf("No cookie cluster")
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantCookieConfig(), fmt.Errorf("The cookie cluster not found")
}

func GetClusterMultiTenantCookieConfigSample() *ClusterMultiTenantCookieConfig {
	c := NewClusterMultiTenantCookieConfig().
		AppendClusters(*GetMultiTenantCookieConfigSample(), *GetMultiTenantCookieConfigSample().SetKey("tenant_2"))
	return c
}
